package domain

import (
	"context"

	pkgerror "github.com/okm321/mahking/go/pkg/error"
)

type Group struct {
	ID      int64     // id
	UID     string    // uuid
	Name    string    // グループ名
	Members []*Member // グループメンバー
	Rule    *Rule     // グループに紐づくルール
}

const (
	MaxGroupNameLength = 100
	MaxGroupMembers    = 10
)

type NewGroupArgs struct {
	Name    string
	Members []*Member
	Rule    *Rule
}

func NewGroup(args NewGroupArgs) (_ *Group, err error) {
	grp := Group{
		Name:    args.Name,
		Members: args.Members,
		Rule:    args.Rule,
	}

	err = grp.Validate()
	if err != nil {
		return nil, err
	}

	return &Group{
		Name:    grp.Name,
		Members: grp.Members,
		Rule:    grp.Rule,
	}, nil
}

func (g *Group) Validate() error {
	if err := requireText("グループ名", g.Name, MaxGroupNameLength); err != nil {
		return err
	}
	if len(g.Members) == 0 {
		return pkgerror.NewClientError("グループメンバーは必須です")
	}
	if len(g.Members) > MaxGroupMembers {
		return pkgerror.NewClientErrorf("グループメンバーは%d人以内で入力してください", MaxGroupMembers)
	}
	if g.Rule == nil {
		return pkgerror.NewClientError("ルールは必須です")
	}
	if required := g.Rule.MahjongType.RequiredMemberCount(); len(g.Members) < required {
		return pkgerror.NewClientErrorf(
			"%sは最低%d人のメンバーが必要です: %d人",
			g.Rule.MahjongType.String(),
			required,
			len(g.Members),
		)
	}
	return nil
}

// GroupRepository 永続化層のインタフェース
type GroupRepository interface {
	List(ctx context.Context) ([]Group, error)
	GetByUUID(ctx context.Context, uid string) (*Group, error)
	Create(ctx context.Context, group *Group) (*Group, error)
}
