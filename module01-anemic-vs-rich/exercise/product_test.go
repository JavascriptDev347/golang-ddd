package exercise

// Bu fayl HOZIRCHA COMPILE BO'LMAYDI — bu me'yor (TDD "qizil holat").
// Sen product.go faylidagi TODO'larni bajarganingdan keyin
// (NewProduct, ReduceStock, Deactivate, getterlar) bu testlar avval
// compile bo'ladi, keyin PASS bo'la boshlaydi.
//
// Ishga tushirish: go test ./module01-anemic-vs-rich/exercise/... -v

// func TestNewProduct_EmptyName_ReturnsError(t *testing.T) {
// 	_, err := NewProduct("p1", "", 1000, 10)
// 	if !errors.Is(err, ErrInvalidName) {
// 		t.Fatalf("expected ErrInvalidName, got %v", err)
// 	}
// }

// func TestNewProduct_InvalidPrice_ReturnsError(t *testing.T) {
// 	_, err := NewProduct("p1", "Telefon", 0, 10)
// 	if !errors.Is(err, ErrInvalidPrice) {
// 		t.Fatalf("expected ErrInvalidPrice, got %v", err)
// 	}
// }

// func TestNewProduct_Valid_CreatesActiveProduct(t *testing.T) {
// 	p, err := NewProduct("p1", "Telefon", 1_500_000, 10)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if !p.IsActive() {
// 		t.Fatalf("new product must be active by default")
// 	}
// 	if p.Stock() != 10 {
// 		t.Fatalf("expected stock 10, got %d", p.Stock())
// 	}
// }

// func TestProduct_ReduceStock_NegativeQty_ReturnsError(t *testing.T) {
// 	p, _ := NewProduct("p1", "Telefon", 1000, 10)
// 	err := p.ReduceStock(-1)
// 	if !errors.Is(err, ErrInvalidQuantity) {
// 		t.Fatalf("expected ErrInvalidQuantity, got %v", err)
// 	}
// }

// func TestProduct_ReduceStock_MoreThanAvailable_ReturnsError(t *testing.T) {
// 	p, _ := NewProduct("p1", "Telefon", 1000, 5)
// 	err := p.ReduceStock(10)
// 	if !errors.Is(err, ErrInsufficientStock) {
// 		t.Fatalf("expected ErrInsufficientStock, got %v", err)
// 	}
// 	// invariant: xato bo'lganda stock o'zgarmasligi kerak
// 	if p.Stock() != 5 {
// 		t.Fatalf("stock should remain 5 after failed reduction, got %d", p.Stock())
// 	}
// }

// func TestProduct_ReduceStock_Success(t *testing.T) {
// 	p, _ := NewProduct("p1", "Telefon", 1000, 5)
// 	if err := p.ReduceStock(3); err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if p.Stock() != 2 {
// 		t.Fatalf("expected stock 2, got %d", p.Stock())
// 	}
// }

// func TestProduct_ReduceStock_WhenInactive_ReturnsError(t *testing.T) {
// 	p, _ := NewProduct("p1", "Telefon", 1000, 5)
// 	_ = p.Deactivate()

// 	err := p.ReduceStock(1)
// 	if !errors.Is(err, ErrProductInactive) {
// 		t.Fatalf("expected ErrProductInactive, got %v", err)
// 	}
// }

// func TestProduct_Deactivate_SetsInactive(t *testing.T) {
// 	p, _ := NewProduct("p1", "Telefon", 1000, 5)
// 	_ = p.Deactivate()
// 	if p.IsActive() {
// 		t.Fatalf("expected product to be inactive after Deactivate()")
// 	}
// }
