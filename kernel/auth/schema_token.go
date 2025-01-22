package auth

func NewVisitor() (ISessId, error) {
	return &tSessId{mode: BROWSE_MODE_VISITOR}, nil
}

func NewPixiv(sessionId string) (ISessId, error) {
	return newToken(BROWSE_MODE_PIXIV, sessionId)
}

func NewFanbox(sessionId string) (ISessId, error) {
	return newToken(BROWSE_MODE_FANBOX, sessionId)
}
