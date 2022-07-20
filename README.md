Projekat iz predmeta Napredne tehnike programiranja

Aplikacija za pronalazak/postavljanje oglasa za posao.

Neregistrovani korisnik:
  - pregled poslova
  - pregled poslodavaca
  - pregled kompanija
  - registracija kao radnik
  - podnošenje zahteva za registraciju kao poslodavac
  - podnošenje zahteva za registraciju kao administrator kompanije

Radnik:
  - prijava na sistem
  - odjava sa sistema
  - izmena profila
  - pregled, filtriranje i sortiranje poslova po određenim parametrima:ocena poslodavca, plata od/do, tip posla(fizički posao, programiranje, uslužna delatnost...), vrsta ugovora(jednokratni, part-time, full-time..), kompanija
  - pregled, filtriranje i sortiranje poslodavaca po određenim parametrima: prosečna plata, broj zaposlenih, prosečna ocena..
  - pregled, filtriranje i sortiranje kompanija po određenim parametrima: prosečna plata, broj zaposlenih, prosečna ocena..
  - detaljan pregled oglasa za posao 
  - prijava na oglas 
  - ocenjivanje i komentarisanje poslodavaca/kompanija u koliko su imali saradnju u prošlosti
  
Samostalni poslodavac:
  - prijava na sistem
  - odjava sa sistema
  - izmena profila
  - crud ponuda za posao (definisanje plate, lokacije, tipa posla, vrste ugovora..)
  - pregled/prihvatanje/odbijanje prijava ranika na oglase koje je postavio
  - ocenjivanje i komentarisanje poslodavaca/kompanija u koliko su imali saradnju u prošlosti
  
Poslodavac kompanije:
  - sve funkcionalnosti samostalnog poslodavca
  - sve ponude koje pravi vezane su za kompaniju koja ga zapošljava
  - pregled/prihvatanje/odbijanje prijava radnika na oglase koje je postavio bilo koji poslodavac iste kompanije
  
Administrator kompanije:
  - prijava na sistem
  - odjava sa sistema
  - izmena profila
  - popunjavanje/izmena profila kompanije
  - crud profila poslodavaca kompanije
  
Administrator sistema:
  - prihvatanje/odbijanje prijava zahteva za registraciju
  - brisanje kompanija i poslodavaca

Servisi:
  - korisnički servis (Go/PostgrateSQL)
  - servis za komentare (Go/PostgrateSQL)
  - servis za ocenjivanje (Go/PostgrateSQL)
  - servis za oglase (Go/PostgrateSQL)
  - servis za kompanije (Go/PostgrateSQL)

Klijentska veb aplikacija:
  - Monolitna angular aplikacija
