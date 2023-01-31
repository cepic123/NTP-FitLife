use actix_web::{web::Data, get, App, HttpServer};
use actix_cors::Cors;
use serde::{Deserialize, Serialize};
use sqlx::{postgres::PgPoolOptions, Pool, Postgres, FromRow}; 

mod block;
use block::services;

pub struct AppState {
    db: Pool<Postgres>
}

#[derive(Serialize, Deserialize, Clone, FromRow)]
struct Block {
    id: i32,
    user_id: i32,
    block_subject_id: i32,
    user_name: String,
    subject_name: String,
}

#[get("/")]
async fn index() -> String {
    "This is a heahth check".to_string()
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let db_url = "postgres://postgres:root@localhost/blockDB";
    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&db_url)
        .await
        .expect("Error building a connection pool");

    HttpServer::new(move || {
        let cors = Cors::default()
              .allowed_origin("localhost:3000")
              .block_on_origin_mismatch(true);

        App::new()
            .wrap(cors)
            .app_data(Data::new(AppState { db: pool.clone()}))
            .service(index)
            .configure(services::config)
    })
    .bind(("localhost", 4002))?
    .run()
    .await
}