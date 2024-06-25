# Simple Risk Assessment API
This API will consist of CRUD user and CRUD submission API. It will be able to calculate risk profile data per user through submission of answer in questionnaire.
Form and calculation for risk profile will be based on this [questionnaire](https://www.danamon.co.id/-/media/ALL-CONTENT-PERSONAL-BANKING/PERSONAL-PRODUCT/PRODUK/INVESTMENT/PDF-File/1/Formulir-Profil-Risiko-Nomor-Referensi-FI-FPR-RISK-0118-009-VERSI-20.pdf?la=id&hash=7472BE2C3273477F0F55E6F5E0B12F290E8C8CB5)

# How To Run
    go run main.go

# Database Migration
## Prerequisite
Install https://github.com/golang-migrate/migrate in local
## Execute
```
migrate -database postgresql://postgres:postgres@localhost:5432/postgres -path assignment-1-risk-assessment/migrations up
```
## Rollback
```
migrate -database postgresql://postgres:postgres@localhost:5432/postgres -path assignment-1-risk-assessment/migrations down
```

## If have dirty version error
```
migrate -database postgresql://postgres:postgres@localhost:5432/postgres -path assignment-1-risk-assessment/migrations force 000001
```