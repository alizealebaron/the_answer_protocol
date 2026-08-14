use std::net::SocketAddr;
use tokio::io::{AsyncBufReadExt, AsyncWriteExt, BufReader};
use tokio::net::TcpStream;

enum Commands {
    // WHERE,
    MOVE,
    LOOK,
    // SEARCH,
    TAKE,
    DROP,
    INVENTORY,
    STATUS,
    ATTACK,
    TALK,
    QUEST,
    QUESTS,
    GROUP,
    CONNECT,
    QUIT,
    CHAT,
    WHO,
    // HELP
}

pub async fn client_manager(mut socket: TcpStream, addr: SocketAddr)
{
    let (reader, mut writer) = socket.split();
    let mut lines = BufReader::new(reader).lines();

    let mut message = format!("blabla fais la commande 'CONNECT'");
    let _ = writer.write_all(message.as_bytes()).await;

    let mut connected: bool = false;

    while let Ok(Some(line)) = lines.next_line().await
    {
        let line: Vec<&str> = line.split(" ").collect();
        if connected == false
        {
            if line[0].trim() == "CONNECT" && line.len() == 3
            {
                connected = true;
                message = format!("Congraulation {}, you are log in", line[1]);
                let _ = writer.write_all(message.as_bytes()).await;
            }
            else if line[0].trim() == "HELP" && line.len() == 1 || line.len() == 0
            {
                continue;
            }
            else
            {
                message = format!("Use the following command'CONNECT [Name] [Language]'");
                let _ = writer.write_all(message.as_bytes()).await;
            }
        }
        else
        {
            continue;
        }
    }

    println!("Client disconnected : {}", addr);
}
