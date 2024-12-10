package main

import(
    "fmt"
    "log"
    "os"
    "os/exec"
)

const (
    VIDEODIR = "vid_in"
    AUDIODIR = "audio_out"
    VIDIN = "/dumbliveremixes8_12102024.mp4"
    OUTNAME = "/dumbliveremixes8_12102024.mp3"
)

func checkCreateDir(dirName string) error {
    info, err := os.Stat(dirName)
    if os.IsNotExist(err) {
        err := os.MkdirAll(dirName, 0755)
        if err != nil {
            return fmt.Errorf("main.go : checkCreateDir() :: ERROR ::: Failed to create directory %s", dirName)
        }
        fmt.Println("main.go : CheckCreateDir() :: Created directory %s", dirName)
    } else if err != nil {
        return fmt.Errorf("main.go : checkCreateDir() :: ERROR ::: An error occured while checking directory %s", dirName)
    } else if !info.IsDir() {
        return fmt.Errorf("main.go : checkCreateDir() :: ERROR :: Path %s exists but it is not a directory", dirName)
    } else {
        fmt.Printf("\nmain.go : checkCreateDir() :: Directory %s already exisits\n", dirName)
    }
    return nil
}

func main() {
    checkCreateDir(VIDEODIR)
    checkCreateDir(AUDIODIR)
    ffmpegCmd := exec.Command(
        "ffmpeg",
        "-i", VIDEODIR+VIDIN,
        "-vn", "-acodec", "libmp3lame", "-q:a", "4",
        AUDIODIR+OUTNAME,
    )
    ffmpegOutput, err := ffmpegCmd.CombinedOutput()
    if err != nil {
        log.Fatalf("\nmain(): An error occured while running ffmpegCmd: %s\n%v", string(ffmpegOutput), err)
    }
    fmt.Printf("\nmain(): Successfully extracted audio from %s and saved as %s with ffmpeg\n\n%s", VIDEODIR+VIDIN, AUDIODIR+OUTNAME, string(ffmpegOutput))
}
