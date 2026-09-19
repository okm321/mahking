package domain

import "context"

type Member struct {
	ID      int64
	GroupID int64
	Name    string // 名前
}

type NewMemberArgs struct {
	Name string
}

const MaxMemberNameLength = 10

func (a NewMemberArgs) validate() error {
	return requireText("名前", a.Name, MaxMemberNameLength)
}

func NewMember(groupID int64, args NewMemberArgs) (_ *Member, err error) {
	if err = args.validate(); err != nil {
		return nil, err
	}

	return &Member{
		GroupID: groupID,
		Name:    args.Name,
	}, nil
}

type MemberRepository interface {
	BatchCreateMembers(ctx context.Context, members []*Member) error
}
