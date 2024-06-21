package mukagen

// Generate mockery for unit test
//go:generate mockery --inpackage --dir=internal/controller --all
//go:generate mockery --inpackage --dir=internal/repository --exclude=internal/repository/ormmodel --all
//go:generate mockery --inpackage --dir=internal/blockchain --exclude=internal/blockchain/contracts --all

// Generate abigen for smart contract interactions
//go:generate abigen --abi build/contracts/Attendance.abi --bin build/contracts/Attendance.bin --pkg contracts --type Attendance --out internal/blockchain/contracts/attendance.go
// Add more line for new contract here
