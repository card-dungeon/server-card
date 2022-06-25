package db

import (
	cardpb "github.com/card-dungeon/server-card/protos/v1/card"
)

var CharCardList = []*cardpb.CharacterCardMessage{
	{
		CardId: 0,
		Name:   "제임스",
		Desc:   "평범한 시골 청년",
		Sprite: "james.png",
		Atk:    3,
		Hp:     10,
		Sd:     1,
		CharacterSkill: &cardpb.CharacterSkill{
			Name:     "거지 근성",
			Desc:     "거지라서 뭘 먹어도 체력을 회복한다. 눈 * 1 체력 회복",
			Cooldown: 1,
		},
	},
	{
		CardId: 1,
		Name:   "정선우",
		Desc:   "치킨 남바완",
		Sprite: "sunu.png",
		Atk:    1,
		Hp:     6,
		Sd:     3,
		CharacterSkill: &cardpb.CharacterSkill{
			Name:     "튀김은 어떠신지?",
			Desc:     "가열된 기름에 물 묻은 닭을 넣어버렸다. 눈 * 실드 수치 공격",
			Cooldown: 3,
		},
	},
	{
		CardId: 2,
		Name:   "박진영",
		Desc:   "진정한 옴붙음",
		Sprite: "jinyong.png",
		Atk:    6,
		Hp:     12,
		Sd:     0,
		CharacterSkill: &cardpb.CharacterSkill{
			Name:     "내 불행을 맛보아라!",
			Desc:     "그 누구도 이길 수 없는 악운의 사나이. 공 - 눈 수치 공격",
			Cooldown: 0,
		},
	},
	{
		CardId: 3,
		Name:   "임현도",
		Desc:   "기재부 씹새끼들",
		Sprite: "hyundo.png",
		Atk:    0,
		Hp:     15,
		Sd:     5,
		CharacterSkill: &cardpb.CharacterSkill{
			Name:     "존나 열받네",
			Desc:     "어떤 극딜도 버티는 정신력의 소유자. 눈 * 2 만큼 아군에게 실드 부여",
			Cooldown: 3,
		},
	},
	{
		CardId: 4,
		Name:   "조재호",
		Desc:   "시대의 싸이코",
		Sprite: "jeho.png",
		Atk:    2,
		Hp:     5,
		Sd:     0,
		CharacterSkill: &cardpb.CharacterSkill{
			Name:     "멘탈 어택",
			Desc:     "정신이 몽롱해지는 입담으로 모두의 뇌를 녹여버린다. 눈 * 공 수치 적 전체 공격",
			Cooldown: 4,
		},
	},
	{
		CardId: 5,
		Name:   "김명훈",
		Desc:   "여긴 누구 나는 어디",
		Sprite: "hunhun.png",
		Atk:    1,
		Hp:     15,
		Sd:     0,
		CharacterSkill: &cardpb.CharacterSkill{
			Name:     "안돼",
			Desc:     "정신이 나가버렸다. 아군이 쓰러질 때 마다 특수 효과를 얻는다. (1명 - 공 1 증가 2명 - 공 2 증가 3명 - 자폭(10 수치 전체 공격))",
			Cooldown: 0,
		},
	},
	{
		CardId: 6,
		Name:   "신창하",
		Desc:   "이것이 기계와 나의 눈높이다",
		Sprite: "chanhaa.png",
		Atk:    2,
		Hp:     7,
		Sd:     0,
		CharacterSkill: &cardpb.CharacterSkill{
			Name:     "일장설명",
			Desc:     "당신과 기계의 눈높이를 열심히 설명한다. 눈 * 공 수치 아군 전체 회복",
			Cooldown: 2,
		},
	},
}
