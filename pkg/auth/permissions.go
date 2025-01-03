package auth

import "encoding/json"

// Permissions это набор разрешений пользователя
type Permissions map[Permission]struct{}

// NewPermissions создает новый набор разрешений пользователя
func NewPermissions(permissions ...Permission) *Permissions {
	p := make(Permissions, len(permissions))
	for _, permission := range permissions {
		p[permission] = struct{}{}
	}
	return &p
}

// Add добавляет разрешения в набор
func (p Permissions) Add(permissions ...Permission) {
	for _, permission := range permissions {
		p[permission] = struct{}{}
	}
}

// Has проверяет наличие разрешения в наборе
func (p Permissions) Has(permission Permission) bool {
	_, ok := p[permission]
	return ok
}

// HasAny проверяет наличие хотя бы одного разрешения в наборе
func (p Permissions) HasAny(permissions ...Permission) bool {
	for _, permission := range permissions {
		if p.Has(permission) {
			return true
		}
	}
	return false
}

// HasAll проверяет наличие всех разрешений в наборе
func (p Permissions) HasAll(permissions ...Permission) bool {
	for _, permission := range permissions {
		if !p.Has(permission) {
			return false
		}
	}
	return true
}

// Slice возвращает слайс разрешений
func (p Permissions) Slice() []Permission {
	permissions := make([]Permission, 0, len(p))
	for key := range p {
		permissions = append(permissions, key)
	}
	return permissions
}

// MarshalJSON сериализует набор разрешений в JSON
func (p Permissions) MarshalJSON() ([]byte, error) {
	keys := make([]string, 0, len(p))
	for key := range p {
		keys = append(keys, string(key))
	}
	return json.Marshal(keys)
}

// UnmarshalJSON десериализует набор разрешений из JSON
func (p *Permissions) UnmarshalJSON(data []byte) error {
	var keys []string
	if err := json.Unmarshal(data, &keys); err != nil {
		return err
	}

	*p = make(Permissions, len(keys))
	for _, key := range keys {
		(*p)[Permission(key)] = struct{}{}
	}

	return nil
}
