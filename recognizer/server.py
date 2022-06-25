
import os
from flask import Flask, request
from werkzeug.utils import secure_filename
import json

import detector
recognizer = detector.Recognizer()

UPLOAD_FOLDER = os.path.join(os.getcwd(), 'photos')
ALLOWED_EXTENSIONS = set(['png', 'jpg', 'jpeg'])
PORT = 51267

if not os.path.exists("photos"):
    os.makedirs("photos")

app = Flask(__name__)
app.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER

def allowed_file(filename):
    return '.' in filename and \
           filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS

@app.route("/recognition/emotion", methods=["POST"]) 
def upload_image():
    file = request.files['file']

    if not file or file.filename == '':
        return 'No selected file'

    if not allowed_file(file.filename):
        return 'No selected file'
    
    filename = secure_filename(file.filename)
    filepath = os.path.join(app.config['UPLOAD_FOLDER'], filename)
    file.save(filepath)
    
    faces = recognizer.detect_faces(filepath)
    results = list()

    for face in faces:
        person_emotions = recognizer.emotion(face)
        current_person = recognizer.find(face)
        return json.dumps({"emotions": person_emotions, "person": current_person})

    dump = json.dumps({"persons": results})
    # return dump.replace("\\", "")
    # return ""
        

app.run("localhost", PORT)