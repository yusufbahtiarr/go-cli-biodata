## ERD CLI Biodata

```mermaid
  erDiagram
  direction LR

  profile {
    int id PK
    string name
    int age
    string gender "ENUM('pria', 'wanita')"
    timestamp created_at
  }

```
