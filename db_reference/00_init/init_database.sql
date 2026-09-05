-- =============================================================================
-- NAVAL ENTERPRISE RESOURCE & READINESS MANAGEMENT ARCHITECTURE (NAVALERP)
-- FILE: 00_init/init_database.sql (Reset Schema & Inisialisasi Database)
-- =============================================================================

-- Bersihkan schema public lama untuk re-generate ulang secara bersih
DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA public;

GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO public;

-- Aktifkan ekstensi yang dibutuhkan
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

