import subprocess
import os
Project_path = os.path.realpath(os.path.curdir)
subprocess.call(["git","init"])
subprocess.call(["git","add","*"])
subprocess.call(["git","commit", "-m","Initial commit"])
# subprocess.call(["git","remote","add","origin", "https://"+Project_path])
# subprocess.call(["git","push","-u","origin","master"])
# subprocess.call(["go","mod","init"])
