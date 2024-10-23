package backend_go

import (
	"bytes"
	"fmt"
	"gocv.io/x/gocv" // OpenCV在Go中的实现，用于处理视频
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

func main() {
	// 打开视频文件
	video, err := gocv.VideoCaptureFile("gesture_video.mp4")
	if err != nil {
		fmt.Printf("Error opening video: %v\n", err)
		return
	}
	defer video.Close()

	frame := gocv.NewMat()
	for {
		// 读取视频帧
		if ok := video.Read(&frame); !ok || frame.Empty() {
			break
		}

		// 将帧转为JPEG格式
		buf := new(bytes.Buffer)
		jpegImg, _ := gocv.IMEncode(".jpg", frame)
		jpegFile := bytes.NewReader(jpegImg)

		// 创建multipart表单请求
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("frame", "frame.jpg")
		io.Copy(part, jpegFile)
		writer.Close()

		// 发送请求到Python API
		resp, err := http.Post("http://localhost:5000/predict", writer.FormDataContentType(), body)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		// 读取响应
		result, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Prediction result:", string(result))

		// 等待下一帧
		time.Sleep(time.Millisecond * 30) // 根据帧率调整
	}
}
