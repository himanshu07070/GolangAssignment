# GolangAssignment

Build a chat server that includes the following points:
1: A user can join a chat group. (using groupID)
2: A user can create a chat group.
3: Max 5 users can join the group.
4: If there is no activity for 30 seconds, prompted a message in chat and closed the group after the next 10 seconds.



  
#### Execution
  * Install all the dependencies 
    * Download Telnet (for ex: via homebrew)
      - To connect different clients to your localhost server at port(for eg: 8000) , then type `telnet localhost [portnumber]`
    
  * Clone the repository
  * Move to the folder in command line and open at least 3 terminals to check exactly how this chat server works:
    ```
  First terminal:
    go build . 
    ./chatbot
   
    ```
   Second and Third terminal (Here you can actually chat between two people in a group):
    ```
    telnet localhost 8000
    ```

    ```
    
   #### Help Commands

  1: /username <username> : Enter Name , otherwise you will stay "Annonymas"
  2: /join <groupname> : Join a room, if room doesn't exist , the new room will be created . User can be only in one room at same time.
  3: /listrooms : Show list of available rooms to join
  4: /msg <message> : broadcast message to everyone in a room
  5: /quit : disconnects from chat server
 
