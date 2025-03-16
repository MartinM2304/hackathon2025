# control_service.py (runs alongside FFmpeg)
import requests
import serial
import time
import threading
from playsound import playsound

# Configuration
API_URL = "https://server.g8row.xyz/api/data"
ARDUINO_PORT = "/dev/ttyUSB0"
POLL_INTERVAL = 5  # Seconds
newline = "\n"
sounds = ["bark","meow","hello","bye"]

class RobotController:
    def __init__(self):
#        self.ser = serial.Serial(ARDUINO_PORT, 115200, timeout=1)
        self.last_command = None
        
    def fetch_command(self):
        try:
            response = requests.get(
                API_URL,
                headers={"Content-Type": "application/json"},
                timeout=3  # Fail fast
            )
            return (response.json()['direction'], response.json()['emoji'], response.json()['sound'])
        except Exception as e:
            print(f"API Error: {str(e)}")
            return None

    def execute_command(self, direction, emoji, sound):
            if direction != None:
 #               self.ser.write(str(direction["Id"]).encode()+newline.encode())    
                print(f"{direction}" )
            if emoji != None:
  #              self.ser.write(str(emoji["Id"]).encode()+newline.encode())
                print(f"{emoji}" )
            if sound != None:
                playsound(f"/home/pi/hackathon2025/rpiclient/{sounds[sound['Id']]}.mp3")        
                print(f"{sound}" )

def control_loop():
    controller = RobotController()
    while True:
        start_time = time.monotonic()
        
        command = controller.fetch_command()
        
        controller.execute_command(command[0],command[1],command[2])
                
        # Dynamic interval adjustment
        elapsed = time.monotonic() - start_time
        sleep_time = max(0, POLL_INTERVAL - elapsed)
        time.sleep(sleep_time)

if __name__ == "__main__":
    # Start in separate thread to keep responsive
    threading.Thread(target=control_loop, daemon=True).start()
    
    # Monitor system health
    while True:
        time.sleep(60)
