package services

import (
	"FileManager/src/config"
	"FileManager/src/domain/dtos/files"
	"FileManager/src/domain/models"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// FileService maneja la lógica relacionada con archivos.
type FileService struct {
	client *mongo.Client // Cliente de MongoDB
}

// NewAuthService crea una nueva instancia de AuthService
func NewFilesService(client *mongo.Client) *FileService {
	return &FileService{client: client}
}

// UploadFile maneja la carga de un archivo y lo guarda localmente.
func (fs *FileService) UploadFile(file multipart.File, dto *files.UploadFileDto) ([]interface{}, error) {
	// Crear un nombre único para el archivo usando el ID de usuario y el tiempo actual
	fileName := fmt.Sprintf("%s_%d_%s", dto.UserID.Hex(), time.Now().Unix(), dto.Filename)

	// Definir la ruta donde se almacenará el archivo
	filePath := filepath.Join("uploads", fileName)

	// Crear el directorio si no existe
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return nil, fmt.Errorf("no se pudo crear el directorio: %v", err)
	}

	// Crear el archivo en la ruta definida
	out, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("no se pudo crear el archivo: %v", err)
	}
	defer out.Close()

	// Copiar el contenido del archivo subido al archivo local
	if _, err = io.Copy(out, file); err != nil {
		return nil, fmt.Errorf("error al guardar el archivo: %v", err)
	}

	// Persistir los datos del archivo en la base de datos
	err = fs.saveFileMetadata(dto, filePath)
	if err != nil {
		return nil, fmt.Errorf("error al guardar los metadatos del archivo: %v", err)
	}

	return []interface{}{dto.UserID, dto.File}, nil
}

// saveFileMetadata persiste los metadatos del archivo en MongoDB.
func (fs *FileService) saveFileMetadata(dto *files.UploadFileDto, filePath string) error {
	// Crear el documento File para almacenar en MongoDB
	fileModel := models.File{
		ID:         primitive.NewObjectID(),
		UserID:     dto.UserID,
		FileName:   dto.Filename,
		Path:       filePath,
		MimeType:   dto.MimeType,
		Size:       dto.Size,
		UploadedAt: time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Obtener la colección de archivos
	fileCollection := fs.client.Database(config.LoadEnv().MONGO_BD_NAME).Collection("files")

	// Insertar el archivo en la colección
	_, err := fileCollection.InsertOne(context.Background(), fileModel)
	if err != nil {
		return fmt.Errorf("error al insertar el archivo en MongoDB: %v", err)
	}

	// Actualizar el usuario con el nuevo archivo
	err = fs.updateUserFiles(dto.UserID, fileModel.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar el usuario con el archivo: %v", err)
	}

	return nil
}

// updateUserFiles actualiza el documento del usuario añadiendo el nuevo archivo subido.
func (fs *FileService) updateUserFiles(userID primitive.ObjectID, fileID primitive.ObjectID) error {
	// Obtener la colección de usuarios
	userCollection := fs.client.Database(config.LoadEnv().MONGO_BD_NAME).Collection("users")

	// Primero, obtener el documento del usuario para verificar si ya tiene el campo `files`
	var user bson.M
	err := userCollection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("usuario no encontrado: %v", err)
		}
		return fmt.Errorf("error al buscar el usuario: %v", err)
	}

	// Si el campo `files` no existe o es `nil`, lo inicializamos

	// Si el campo `files` ya existe, simplemente hacer un push
	_, err = userCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{
			"$push": bson.M{"files": fileID}, // Añadir el archivo al array existente
		},
	)
	if err != nil {
		return fmt.Errorf("error al hacer push a los archivos del usuario: %v", err)
	}

	return nil
}

func (fs *FileService) GetAllFiles() ([]models.File, error) {
	// Obtener la colección de archivos
	fileCollection := fs.client.Database(config.LoadEnv().MONGO_BD_NAME).Collection("files")

	cursor, err := fileCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error al obtener archivos: %v", err)
	}
	defer cursor.Close(context.Background())

	// Crear un slice para almacenar los resultados
	var files []models.File

	if err = cursor.All(context.TODO(), &files); err != nil {
		panic(err)
	}

	return files, nil
}
