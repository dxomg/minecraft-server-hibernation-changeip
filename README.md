# minecraft-server-hibernation
concept and early-code by [gekigek99](https://github.com/gekigek99/minecraft-vanilla-server-hibernation)<br/>
contributor (advanced-code) by [najtin](https://github.com/najtin/minecraft-server-hibernation)<br/>
derived from [supernifty](https://github.com/supernifty/port-forwarder)<br/>

This is a simple Python script to start a minecraft server on request and stop it when there are no player online.
How to use:
1. Install and run your desiered minecraft server
2. Rename the minecraft-server-jar to 'minecraft_server.jar'
3. Check the server-port parameter in 'server.properties': must be 25565
4. Edited the paramters in the script as needed. 
5. run the script at reboot
6. you can connect to the server through port 25555

**IMPORTANT**	
If you are the first to access to minecraft world you will *have to wait 30 seconds* and then try to connect again.
```Python
MINECRAFT_SERVER_STARTUPTIME = 30       #any parameter more than 10s is recommended
```
After 120 seconds you have 240 to connect to the server before it is shutdown. 
```Python
TIME_BEFORE_STOPPING_EMPTY_SERVER = 120 #any parameter more than 60s is recommended
```
You can change these parameters to fit your needs.
