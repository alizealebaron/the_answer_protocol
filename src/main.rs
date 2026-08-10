//  import
use tokio::net::TcpListener;
mod client_manager;
use client_manager::client_manager;

//  transformer le main classique synchrone en main assynchrone
#[tokio::main]
//  fonction main async qui revoit rien si ok, ou une erreur si probleme
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("Serveur en cours de démarrage...");
    // ecoute sur le port 4242 (bind) toute les adresse possible
    let listener: TcpListener = TcpListener::bind("0.0.0.0:4343").await?;
    println!("Serveur démarrer");

    loop {
        // attend en arriere plan qu'un client se connect (.accept())
        // on boucle dessus pour recuperer tout les client
        // Pour se connecte les client utilise la commande
        // nc [ip host] [port] -> nc 10.12.12.3 4343
        let (socket, addr) = listener.accept().await?;
        println!("New client connected : {}", addr);

        // Démarrer une tâche indépendante pour ce client
        tokio::spawn(async move {
            // Traitement des commandes du client ici
            client_manager(socket, addr).await;
        });
    }
}
