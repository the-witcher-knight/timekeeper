package mukagen

// Generate mockery for unit test
//go:generate mockery --inpackage --dir=internal/controller --all
//go:generate mockery --inpackage --dir=internal/repository --exclude=internal/repository/ormmodel --all
//go:generate mockery --inpackage --dir=internal/blockchain --exclude=internal/blockchain/contracts --all
