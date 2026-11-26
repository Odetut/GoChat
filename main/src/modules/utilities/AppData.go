package utilities
import ("encoding/json"
        "os"
        "path/filepath"
    
    )


type AppData struct {
    Version string `json:"Version"`
    AppName string `json:"AppName"`
    Release string `json:"Release"`
}

func GetOwnAppData() (AppData,error){
    Path := filepath.Join("..","..","..","..","Metadata","AppInfo.json")
    content,err := os.ReadFile(Path) // Open the file instantly into memory cause its small

    if err != nil{ // Fallback struct to satisfy argument count
        return AppData {Version: "nil" , AppName: "nil" , Release: "nil" }, err
    }
    var data AppData

    if err := json.Unmarshal(content, &data); err != nil {
        return AppData{Version: "nil", AppName: "nil", Release: "nil"}, err
    }
    return data, nil
}


// Handle Extracting data from AppInfo.json and send back to the caller
