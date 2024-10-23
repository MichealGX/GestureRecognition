import cv2
import torch
from flask import Flask, request, jsonify

# 加载YOLOv11模型
model = torch.hub.load('ultralytics/yolov5', 'custom', path='your_yolo_hand_gesture_model.pt')  # 替换为你的YOLOv11模型路径

app = Flask(__name__)

@app.route('/predict', methods=['POST'])
def predict():
    # 从请求中获取视频帧
    file = request.files['frame'].read()
    npimg = np.fromstring(file, np.uint8)
    img = cv2.imdecode(npimg, cv2.IMREAD_COLOR)

    # 使用模型进行预测
    results = model(img)

    # 提取手势识别结果
    labels, cords = results.xyxyn[0][:, -1].numpy(), results.xyxyn[0][:, :-1].numpy()
    predictions = []
    for label, cord in zip(labels, cords):
        predictions.append({
            'label': int(label),
            'coordinates': cord.tolist()
        })

    return jsonify(predictions)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)

cap = cv2.VideoCapture('video_path.mp4')
while cap.isOpened():
    ret, frame = cap.read()
    if not ret:
        break

    # 每一帧可以通过API发送出去处理
    # 这里你可以通过POST请求把frame发送给后端的API
    # 例如上面的 `/predict` API
