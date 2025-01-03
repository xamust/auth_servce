package auth

// OperationPermissions это мапа OpenAPI операций к разрешениям
type OperationPermissions[T comparable] map[T]Permission

// Allowed проверяет разрешение на операцию.
// Возвращает true если операция не защищена или если у пользователя есть разрешение.
func (o OperationPermissions[T]) Allowed(
	operationID T,
	userPermissions *Permissions,
) bool {
	permission, isProtected := o[operationID]
	return !isProtected || userPermissions.Has(permission)
}
