use actix_web::{get, web, App, HttpServer};
use serde::{Deserialize, Serialize};
use std::sync::Mutex;

mod complaint;
use complaint::services;

struct AppState {
    complaints: Mutex<Vec<Complaint>>
}

#[derive(Serialize, Deserialize, Clone)]
struct Complaint {
    id: usize,
    user_id: usize,
    complaint_subject_id: usize,
    complaint_text: String,
}

#[get("/")]
async fn index() -> String {
    "This is a heahth check".to_string()
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let app_data = web::Data::new(AppState {
        complaints: Mutex::new(vec![])
    });

    HttpServer::new(move || {
        App::new()
            .app_data(app_data.clone())
            .service(index)
            .configure(services::config)
    })
    .bind(("localhost", 4001))?
    .run()
    .await
}