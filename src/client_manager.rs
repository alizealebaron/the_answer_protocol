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

pub async fn client_manager(mut socket: TcpStream, addr: SocketAddr) {
    let (reader, mut writer) = socket.split();
    let mut lines = BufReader::new(reader).lines();

    while let Ok(Some(ligne)) = lines.next_line().await {
        let ligne = ligne.trim();
        let cmd_parse = ligne.split(' ').next().unwrap_or("");
        if cmd_list.contains(&cmd_parse) {
            println!("Commande reçue by {}: {}", addr, ligne);
            let response = format!("you sent {}\n", cmd_parse);
            let _ = writer.write_all(response.as_bytes()).await;
        }
        else {
            let response = format!("Unknown command\n");
            let _ = writer.write_all(response.as_bytes()).await;
        }
    }

    println!("Client disconnected : {}", addr);
}
