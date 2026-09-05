CREATE DATABASE ojas153;

\c ojas153;


DROP TYPE IF EXISTS tipe_user CASCADE;
CREATE TYPE tipe_user AS ENUM ('admin', 'super_admin', 'penimbang', 'penerima');
CREATE TABLE "user" (
    id_user SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    nama_user VARCHAR(100) NOT NULL,
    no_telp VARCHAR(100),
    tanggal_lahir DATE,
    alamat TEXT,
    tipe_user tipe_user NOT NULL DEFAULT 'admin',
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);



CREATE TABLE provinsi (
    id_provinsi SERIAL PRIMARY KEY,
    nama_provinsi VARCHAR(100) NOT NULL,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);


CREATE TABLE kota (
    id_kota SERIAL PRIMARY KEY,
    id_provinsi INT,
    nama_kota VARCHAR(100) NOT NULL,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_provinsi) REFERENCES provinsi(id_provinsi),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

    

CREATE TABLE pelabuhan (
    id_pelabuhan SERIAL PRIMARY KEY,
    id_kota INT,
    nama_pelabuhan VARCHAR(100) NOT NULL,
    alamat TEXT,
    jarak NUMERIC(12,3),
    latitude NUMERIC(9,6),
    longitude NUMERIC(9,6),
    place_id VARCHAR(100),
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    foto TEXT,
    FOREIGN KEY (id_kota) REFERENCES kota(id_kota),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE jetty (
    id_jetty SERIAL PRIMARY KEY,
    id_pelabuhan INT NOT NULL,
    nama_jetty VARCHAR(100) NOT NULL,
    panjang_dermaga  NUMERIC(12,3) NOT NULL,
    kedalaman  NUMERIC(12,3) NOT NULL,
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_pelabuhan) REFERENCES pelabuhan(id_pelabuhan),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE tipe_kapal (
    id_tipe_kapal SERIAL PRIMARY KEY,
    nama_tipe_kapal VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE kapal (
    id_kapal SERIAL PRIMARY KEY,
    id_tipe_kapal INT NOT NULL,
    imo_number INT NOT NULL,
    nama_kapal VARCHAR(100) NOT NULL,
    panjang_kapal NUMERIC(12,3),
    lebar_kapal NUMERIC(12,3),
    kapasitas  NUMERIC(12,3), 
    tahun_pembuatan INT,
    url_kapal VARCHAR(500),
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    foto TEXT,
    FOREIGN KEY (id_tipe_kapal) REFERENCES tipe_kapal(id_tipe_kapal),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE pbm (
    id_pbm SERIAL PRIMARY KEY,
    id_kota INT,
    nama_pbm VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    no_telp VARCHAR(100),
    alamat TEXT,
    jarak NUMERIC(12,3),
    latitude NUMERIC(9,6),
    longitude NUMERIC(9,6),
    place_id VARCHAR(100),
    deskripsi TEXT,
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_kota) REFERENCES kota(id_kota),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

DROP TYPE IF EXISTS status_kapal_unloading CASCADE;
CREATE TYPE status_kapal_unloading AS ENUM ('belum_sampai', 'unloading', 'selesai');
CREATE TABLE kapal_unloading (
    id_kapal_unloading SERIAL PRIMARY KEY,
    id_kapal INT NOT NULL,
    id_jetty INT NOT NULL,
    id_pbm INT NOT NULL,
    waktu_sandar TIMESTAMP WITH TIME ZONE NOT NULL,
    waktu_berangkat TIMESTAMP WITH TIME ZONE,
    muatan_total NUMERIC(12,3),
    target_jumlah_siklus_harian int,
    target_jumlah_truk_harian int,
    target_berat_per_truk NUMERIC(12,3), 
    target_berat_per_meter_kubik NUMERIC(12,3), 
    target_durasi_siklus NUMERIC(12,3), 
    target_durasi_timbang NUMERIC(12,3),
    target_durasi_pengangkutan NUMERIC(12,3),
    status_kapal_unloading status_kapal_unloading NOT NULL DEFAULT 'belum_sampai',
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_kapal) REFERENCES kapal(id_kapal),
    FOREIGN KEY (id_jetty) REFERENCES jetty(id_jetty),
    FOREIGN KEY (id_pbm) REFERENCES pbm(id_pbm),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE tipe_truk (
    id_tipe_truk SERIAL PRIMARY KEY,
    nama_tipe_truk VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);


CREATE TABLE sopir (
    id_sopir SERIAL PRIMARY KEY,
    nama_sopir VARCHAR(100) NOT NULL,
    no_telp VARCHAR(100),
    tanggal_lahir DATE,
    alamat TEXT,
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE truk (
    id_truk SERIAL PRIMARY KEY,
    id_tipe_truk INT NOT NULL,
    id_sopir INT NOT NULL,
    nomor_polisi VARCHAR(20) UNIQUE NOT NULL,
    panjang INT,
    lebar INT,
    tinggi INT,
    berat_kosong NUMERIC(12,3), 
    barcode VARCHAR(100),
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_tipe_truk) REFERENCES tipe_truk(id_tipe_truk),
    FOREIGN KEY (id_sopir) REFERENCES sopir(id_sopir),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);


CREATE TABLE tipe_muatan (
    id_tipe_muatan SERIAL PRIMARY KEY,
    nama_tipe_muatan VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);




CREATE TABLE client (
    id_client SERIAL PRIMARY KEY,
    id_kota INT,
    nama_client VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    no_telp VARCHAR(100),
    alamat TEXT,
    jarak NUMERIC(12,3),
    latitude NUMERIC(9,6),
    longitude NUMERIC(9,6),
    place_id VARCHAR(100),
    deskripsi TEXT,
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_kota) REFERENCES kota(id_kota),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

DROP TYPE IF EXISTS tipe_pengangkutan CASCADE;
DROP TYPE IF EXISTS tipe_proyek CASCADE;
CREATE TYPE tipe_pengangkutan AS ENUM ('tonase', 'ritase');
CREATE TYPE status_proyek AS ENUM ('belum_dimulai', 'proses', 'selesai');

CREATE TABLE proyek_client (
    id_proyek_client SERIAL PRIMARY KEY,
    id_client INT NOT NULL,
    id_kota INT,
    nama_proyek VARCHAR(100) NOT NULL,
    alamat TEXT,
    jarak NUMERIC(12,3),
    latitude NUMERIC(9,6),
    longitude NUMERIC(9,6),
    place_id VARCHAR(100),
    deskripsi TEXT,
    berat_total NUMERIC(12,3),
    tipe_pengangkutan tipe_pengangkutan NOT NULL DEFAULT 'ritase',
    status_proyek status_proyek NOT NULL DEFAULT 'belum_dimulai',
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_client) REFERENCES client(id_client),
    FOREIGN KEY (id_kota) REFERENCES kota(id_kota),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);



CREATE TABLE muatan (
    id_muatan SERIAL PRIMARY KEY,
    id_tipe_muatan INT NOT NULL,
    id_proyek_client INT NOT NULL,
    id_kapal_unloading INT NOT NULL,
    nama_muatan VARCHAR(100) NOT NULL,
    berat_total NUMERIC(12,3) NOT NULL,
    massa_jenis NUMERIC(12,3) NOT NULL,
    deskripsi TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_tipe_muatan) REFERENCES tipe_muatan(id_tipe_muatan),
    FOREIGN KEY (id_proyek_client) REFERENCES proyek_client(id_proyek_client),
    FOREIGN KEY (id_kapal_unloading) REFERENCES kapal_unloading(id_kapal_unloading),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- DROP TYPE IF EXISTS tipe_pengangkutan CASCADE;
-- CREATE TYPE tipe_pengangkutan AS ENUM ('ritase', 'tonase');
-- CREATE TABLE proyek_client_muatan (
--     id_proyek_client_muatan SERIAL PRIMARY KEY,
--     id_proyek_client INT NOT NULL,
--     id_muatan INT NOT NULL,
--     tipe_pengangkutan tipe_pengangkutan NOT NULL DEFAULT 'tonase',
--     is_show BOOLEAN NOT NULL DEFAULT TRUE,
--     FOREIGN KEY (id_muatan) REFERENCES muatan(id_muatan),
--     FOREIGN KEY (id_proyek_client) REFERENCES proyek_client(id_proyek_client),
--     created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
--     updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
--     deleted_at TIMESTAMP WITH TIME ZONE
-- );


-- INI KEBAWAH BELOM
DROP TYPE IF EXISTS status_pengangkutan_bongkar CASCADE;
CREATE TYPE status_pengangkutan_bongkar AS ENUM ('akan_mengangkut', 'dalam_perjalanan', 'menurunkan_muatan', 'selesai');
CREATE TABLE pengangkutan_bongkar(
    id_pengangkutan_bongkar SERIAL PRIMARY KEY,
    id_muatan INT NOT NULL,
    id_truk INT NOT NULL,
    id_sopir INT NOT NULL,
    nomor_surat_jalan VARCHAR(100),
    berat_kosong NUMERIC(12,3),
    berat_isi NUMERIC(12,3),
    waktu_jetty TIMESTAMP WITH TIME ZONE NOT NULL,
    waktu_timbang_isi TIMESTAMP WITH TIME ZONE,
    waktu_penurunan_muatan TIMESTAMP WITH TIME ZONE,
    status_pengangkutan_bongkar status_pengangkutan_bongkar NOT NULL DEFAULT 'akan_mengangkut',
    foto_kosong VARCHAR(200),
    foto_isi VARCHAR(200),
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_truk) REFERENCES truk(id_truk),
    FOREIGN KEY (id_sopir) REFERENCES sopir(id_sopir),
    FOREIGN KEY (id_muatan) REFERENCES muatan(id_muatan),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE tipe_quarry (
    id_tipe_quarry SERIAL PRIMARY KEY,
    nama_tipe_quarry VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    foto TEXT,  
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE quarry (
    id_quarry SERIAL PRIMARY KEY,
    id_tipe_quarry INT NOT NULL,
    id_kota INT,
    iup VARCHAR(100)  NOT NULL,
    iup_op VARCHAR(100)  NOT NULL,
    nama_quarry VARCHAR(100) NOT NULL,
    alamat TEXT,
    jarak NUMERIC(12,3),
    latitude NUMERIC(9,6),
    longitude NUMERIC(9,6),
    place_id VARCHAR(100),
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_tipe_quarry) REFERENCES tipe_quarry(id_tipe_quarry),
    FOREIGN KEY (id_kota) REFERENCES kota(id_kota),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE stockpile (
    id_stockpile SERIAL PRIMARY KEY,
    id_quarry INT NOT NULL,
    id_kota INT,
    nama_stockpile VARCHAR(100) NOT NULL,
    alamat TEXT,
    jarak NUMERIC(12,3),
    latitude NUMERIC(9,6),
    longitude NUMERIC(9,6),
    place_id VARCHAR(100),
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_quarry) REFERENCES quarry(id_quarry),
    FOREIGN KEY (id_kota) REFERENCES kota(id_kota),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);


CREATE TABLE tipe_material (
    id_tipe_material SERIAL PRIMARY KEY,
    nama_tipe_material VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    foto TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE material (
    id_material SERIAL PRIMARY KEY,
    id_tipe_material INT NOT NULL,
    id_proyek_client INT NOT NULL,
    id_stockpile INT NOT NULL,
    nama_material VARCHAR(100) NOT NULL,
    berat_total NUMERIC(12,3),
    massa_jenis NUMERIC(12,3),
    deskripsi TEXT,
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_tipe_material) REFERENCES tipe_material(id_tipe_material),
    FOREIGN KEY (id_proyek_client) REFERENCES proyek_client(id_proyek_client),
    FOREIGN KEY (id_stockpile) REFERENCES stockpile(id_stockpile),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);


DROP TYPE IF EXISTS status_pengangkutan_muat CASCADE;
CREATE TYPE status_pengangkutan_muat AS ENUM ('akan_mengangkut', 'dalam_perjalanan', 'menurunkan_muatan', 'selesai');
CREATE TABLE pengangkutan_muat(
    id_pengangkutan_muat SERIAL PRIMARY KEY,
    id_material INT NOT NULL,
    id_truk INT NOT NULL,
    id_sopir INT NOT NULL,
    nomor_surat_jalan VARCHAR(100),
    berat_kosong NUMERIC(12,3),
    berat_isi NUMERIC(12,3),
    waktu_tambang TIMESTAMP WITH TIME ZONE NOT NULL,
    waktu_timbang_isi TIMESTAMP WITH TIME ZONE,
    waktu_penurunan_muatan TIMESTAMP WITH TIME ZONE,
    status_pengangkutan_muat status_pengangkutan_muat NOT NULL DEFAULT 'akan_mengangkut',
    foto_berangkat VARCHAR(200),
    foto_pulang VARCHAR(200),
    is_show BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (id_truk) REFERENCES truk(id_truk),
    FOREIGN KEY (id_sopir) REFERENCES sopir(id_sopir),
    FOREIGN KEY (id_material) REFERENCES material(id_material),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);



stockpile
tipe_material
client
quarry
pengangkutan_muat


untuk barcode scan
keluar stockpile , sesudah timbang isi, input jumlah muatan

keluar pabrik, input jumlah muatan informasi dari pabrik


 




-- DROP TYPE IF EXISTS status_pengangkutan_tonase CASCADE;
-- CREATE TYPE status_pengangkutan_tonase AS ENUM ('belum_mengangkut', 'sedang_mengangkut', 'dalam_perjalanan', 'menurunkan_muatan', 'selesai');
-- CREATE TABLE pengangkutan_tonase (
--     id_pengangkutan_tonase SERIAL PRIMARY KEY,
--     id_proyek_client_muatan INT NOT NULL,
--     id_truk INT NOT NULL,
--     nomor_surat_jalan VARCHAR(100),
--     berat_kosong NUMERIC(12,3),
--     berat_kotor NUMERIC(12,3),
--     waktu_timbang_kosong TIMESTAMP WITH TIME ZONE NOT NULL,
--     waktu_timbang_kotor TIMESTAMP WITH TIME ZONE,
--     waktu_penurunan_muatan TIMESTAMP WITH TIME ZONE,
--     status_pengangkutan_tonase status_pengangkutan_tonase NOT NULL DEFAULT 'belum_mengangkut',
--     FOREIGN KEY (id_truk) REFERENCES truk(id_truk),
--     FOREIGN KEY (id_proyek_client_muatan) REFERENCES proyek_client_muatan(id_proyek_client_muatan),
--     created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
--     updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
--     deleted_at TIMESTAMP WITH TIME ZONE
-- );



-- DROP TYPE IF EXISTS tipe_kecurangan CASCADE;
-- CREATE TYPE tipe_kecurangan AS ENUM ('timbang_kosong_tidak_ada', 'timbang_kotor_tidak_ada', 'timbang_lebih_dari_satu', 'timbang_kosong_kotor_terpaut_jauh', 'durasi_satu_siklus','jumlah_siklus_harian', 'jumlah_truk_harian','selesai');

-- CREATE TABLE dugaan_kecurangan (
--     id_dugaan_kecurangan SERIAL PRIMARY KEY,
--     id_kapal_unloading INT NOT NULL,
--     id_truk INT,
--     id_pengangkutan INT,
--     judul VARCHAR(100),
--     pesan VARCHAR(100),
--     tipe_kecurangan tipe_kecurangan NOT NULL DEFAULT 'timbang_kosong_tidak_ada',
--     FOREIGN KEY (id_kapal_unloading) REFERENCES kapal_unloading(id_kapal_unloading),
--     FOREIGN KEY (id_truk) REFERENCES truk(id_truk),
--     FOREIGN KEY (id_pengangkutan) REFERENCES pengangkutan(id_pengangkutan),
--     created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
--     updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
--     deleted_at TIMESTAMP WITH TIME ZONE
-- );












