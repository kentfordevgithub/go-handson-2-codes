// godocを生成するためのパッケージコメント
package profile1

// ユーザープロフィール構造体
type Profile struct {
	ID   int
	Name string
}

// ユーザープロフィールを作成します
func CreateProfile(id int, name string) *Profile {
	profile := &Profile{
		ID:   id,
		Name: name,
	}
	// TODO: profile作成

	return profile
}
