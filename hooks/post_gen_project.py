import subprocess
import os
Project_path = os.path.realpath(os.path.curdir)
project_name='{{cookiecutter.project_name}}'
git_host = '{{cookiecutter.git_hoster}}'
owner = '{{cookiecutter.owner}}'
git_name = '{{cookiecutter.git_name}}'
subprocess.call(["git","init"])
subprocess.call(["git","remote","add","origin", "https://"+git_host+"/"+owner+"/"+project_name+".git"])
subprocess.call(["git","add","."])
subprocess.call(["git","commit", "-m","Initial commit"])
subprocess.call(["git","push","-u","origin","master"])
subprocess.call(["go","mod","init"])
subprocess.call(["go","get"])
