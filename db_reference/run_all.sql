-- =============================================================================
-- NAVAL ENTERPRISE RESOURCE PLANNING (NAVALERP)
-- MASTER DATABASE INITIALIZATION & SEEDING SCRIPT
-- =============================================================================
-- Arsitektur Berdasarkan Dokumen Resmi:
-- "Naval Enterprise Resource Planning Architecture" (Laksamana Pertama TNI Judijanto)
--
-- Eksekusi melalui psql:
-- psql -h 72.62.122.31 -p 5432 -U dutakasih -d navalerp -f run_all.sql
-- =============================================================================

\echo '>>> [00/10] Inisialisasi Database, Schema & Ekstensi...'
\i 00_init/init_database.sql

\echo '>>> [01/10] Modul 1: Organisasi Komando & Pengguna Sistem (1 User 1 Role)...'
\i 01_organization_user/tables.sql
\i 01_organization_user/insert.sql

\echo '>>> [02/10] Modul 2: Manajemen Armada Kapal Perang, Aset & MRO Terpadu...'
\i 02_fleet_asset_mro/tables.sql
\i 02_fleet_asset_mro/insert.sql

\echo '>>> [03/10] Modul 3: Pergudangan, Suku Cadang & Inventaris Alutsista...'
\i 03_inventory_warehousing/tables.sql
\i 03_inventory_warehousing/insert.sql

\echo '>>> [04/10] Modul 4: Pengadaan Pertahanan & Lelang Industri Alutsista...'
\i 04_defence_procurement/tables.sql
\i 04_defence_procurement/insert.sql

\echo '>>> [05/10] Modul 5: Keuangan Militer, DIPA, Akuntansi & Total Cost of Ownership (TCO)...'
\i 05_finance_budget_accounting/tables.sql
\i 05_finance_budget_accounting/insert.sql

\echo '>>> [06/10] Modul 6: Sumber Daya Manusia Militer & Pengawakan Kapal Perang (HCM)...'
\i 06_military_human_capital/tables.sql
\i 06_military_human_capital/insert.sql

\echo '>>> [07/10] Modul 7: Infrastruktur Pangkalan, Fasilitas Labuh & Sandar Kapal...'
\i 07_base_infrastructure/tables.sql
\i 07_base_infrastructure/insert.sql

\echo '>>> [08/10] Modul 8: Distribusi Logistik Laut Militer & Transportasi Bebekal...'
\i 08_logistics_transportation/tables.sql
\i 08_logistics_transportation/insert.sql

\echo '>>> [09/10] Modul 9: EDRMS, Dokumen Teknis Militer & Digital Thread Alutsista...'
\i 09_edrms_documents/tables.sql
\i 09_edrms_documents/insert.sql

\echo '>>> [10/11] Modul 10: Kesiapan Tempur Alutsista & Komando Operasi Armada (C2)...'
\i 10_readiness_operations_command/tables.sql
\i 10_readiness_operations_command/insert.sql

\echo '>>> [11/11] Modul 11: Infrastruktur Bawah Laut Kritis (Critical Underwater Infrastructure - CUI)...'
\i 11_cui_underwater_infrastructure/tables.sql
\i 11_cui_underwater_infrastructure/insert.sql

\echo '============================================================================='
\echo '>>> SELURUH STRUKTUR DATABASE (65+ TABEL) DAN SEEDER NAVALERP BERHASIL DILAKSANAKAN!'
\echo '============================================================================='

