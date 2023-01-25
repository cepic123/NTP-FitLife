use actix_web::{web::Data, get, web, App, HttpServer};
use serde::{Deserialize, Serialize};
use std::sync::Mutex;
use sqlx::{postgres::PgPoolOptions, Pool, Postgres, FromRow}; 

mod complaint;
use complaint::services;

pub struct AppState {
    db: Pool<Postgres>
}

#[derive(Serialize, Deserialize, Clone, FromRow)]
struct Complaint {
    id: i32,
    user_id: i32,
    complaint_subject_id: i32,
    complaint_text: String,
    user_name: String,
    subject_name: String,
}

#[get("/")]
async fn index() -> String {
    "This is a heahth check".to_string()
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let db_url = "postgres://postgres:root@localhost/complaintDB";
    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&db_url)
        .await
        .expect("Error building a connection pool");

    HttpServer::new(move || {
        App::new()
            .app_data(Data::new(AppState { db: pool.clone()}))
            .service(index)
            .configure(services::config)
    })
    .bind(("localhost", 4001))?
    .run()
    .await
}