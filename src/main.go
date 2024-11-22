package main

import(
  "fmt"
  "log"
  "os/exec"
)

const (
  VIDIN = "src/vid_in/dumbliveremixes2_11222024.mp4"
  OUTNAME = "src/audio_out/dumbliveremixes2_11222024.mp3"
)

func main() {
  ffmpegCmd := exec.Command(
    "ffmpeg",
    "-i", VIDIN,
    "-vn", "-acodec", "libmp3lame", "-q:a", "4",
    OUTNAME,
  )
  ffmpegOutput, err := ffmpegCmd.CombinedOutput()
  if err != nil {
    log.Fatalf("\nmain(): An error occured while running ffmpegCmd: %s\n%v", string(ffmpegOutput), err)
  }
  fmt.Printf("\nmain(): Successfully extracted audio from %s and saved as %s with ffmpeg\n\n%s", VIDIN, OUTNAME, string(ffmpegOutput))
}
