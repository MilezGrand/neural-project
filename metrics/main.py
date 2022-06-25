from this import s
from deepface import DeepFace
from insightface.app import FaceAnalysis

import os
import cv2
import numpy as np
import time

def get_all_files(folder):
    file_list = []
    if os.path.exists(folder):
        for root, dirs, files in os.walk(folder):
            for file in files:
                file_list.append(os.path.join(root,file))
    return file_list

class Recognizer:
    def __init__(self):
        self.__app = FaceAnalysis(providers=['CUDAExecutionProvider', 'CPUExecutionProvider'])
        self.__app.prepare(ctx_id=0, det_size=(640, 640))

        self.PHOTOS_FOLDER = os.path.join(os.getcwd(), 'photos')
        self.PERSONS_FOLDER = os.path.join(os.getcwd(), 'persons')
        self.__THRESHOLD = 0.18

    # Определяет эмоцию человека
    def emotion(self, img_path):
        result = DeepFace.analyze(img_path = img_path, actions = ['emotion'], enforce_detection=False)
        emotions = result["emotion"]
        return emotions

    # Принадлежит ли лицо img_path1 человеку img_path2
    def __verify(self, img_path1, img_path2):
        result = DeepFace.verify(img1_path = img_path1, img2_path = img_path2, enforce_detection=False)
        return float(result["distance"])

    # Ищет лица на фото
    def detect_faces(self, img_path):
        img = cv2.imread(img_path)

        if img is None:
            return []

        faces = self.__app.get(img)
        filenames = list()

        for face in faces:
            box = face.bbox.astype(np.int)

            x0 = box[0]
            y0 = box[1]
            x1 = box[2]
            y1 = box[3]

            cropped_image = img[y0: y1, x0: x1]
            filename = str(round(time.time() * 1000)) + ".jpg"
            filenames.append(filename)

            cv2.imwrite(filename, cropped_image)

        return filenames

    # Получает список людей из папки persons
    def __get_persons(self):
        persons = list()

        for file in os.listdir(self.PERSONS_FOLDER):
            d = os.path.join(self.PERSONS_FOLDER, file)
            if os.path.isdir(d):
                persons.append(d)

        return persons

    # Ищет кому принадлежит лицо
    def find(self, img_path):
        minimal_dist = 1.0
        persons = self.__get_persons()
        current_person = ""

        for person in persons:
            images = get_all_files(person)
            for img in images:
                faces = self.detect_faces(img)
                for face in faces:
                    result = self.__verify(img_path, face)
                    os.remove(face)
                    if result < minimal_dist:
                        minimal_dist = result
                        current_person = img.split("\\")[-2]
        
        return current_person

def __get_emotion_by_index(index):
    if index == 0:
        return "angry"
    elif index == 1:
        return "disgust"
    elif index == 2:
        return "fear"
    elif index == 3:
        return "happy"
    elif index == 4:
        return "sad"
    elif index == 5:
        return "surprise"
    elif index == 6:
        return "neutral"

    return "unknown"

def __get_max_emotion(obj):
    emotions = [float(obj["angry"]), float(obj["disgust"]), float(obj["fear"]), float(obj["happy"]), float(obj["sad"]), float(obj["surprise"]), float(obj["neutral"])]
    maximal = max(emotions)
    for i in range(0, len(emotions)):
        if emotions[i] == maximal:
            return i

if __name__ == "__main__":
    recognizer = Recognizer()
    emotions = os.listdir("emotions")
    s = 0

    current_time = time.time()

    for emotion in emotions:
        faces = recognizer.detect_faces(os.path.join(os.getcwd(), 'emotions', emotion))
        for face in faces:
            emotions_of_person = recognizer.emotion(face)
            index = __get_max_emotion(emotions_of_person)
            emotion_of_person = __get_emotion_by_index(index)
            print(emotion_of_person, emotion.split("_")[0])
            if emotion_of_person == "surprise":
                if emotion.split("_")[0] == "fear" or emotion.split("_")[0] == "surprise":
                    s += 1    
            elif emotion_of_person == emotion.split("_")[0]:
                s += 1
            os.remove(face)

    print(time.time() - current_time)
    print(f'{s} / {len(emotions)}', s / len(emotions))