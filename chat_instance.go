package dalgo4botsfwtelegram

import (
	"context"
	telegram "github.com/bots-go-framework/bots-fw-telegram"
	"github.com/bots-go-framework/bots-fw-telegram-models/botsfwtgmodels"
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo/record"
)

type TgChatInstance struct {
	record.WithID[string]
	Data botsfwtgmodels.TgChatInstanceData
}

var _ telegram.TgChatInstanceDal = (*tgChatInstanceDalgo)(nil)

type tgChatInstanceDalgo struct {
	db dal.Database
}

func (tgChatInstanceDal tgChatInstanceDalgo) GetTelegramChatInstanceByID(c context.Context /*tx dal.ReadTransaction,*/, id string) (tgChatInstanceData botsfwtgmodels.TgChatInstanceData, err error) {
	tgChatInstanceData = tgChatInstanceDal.NewTelegramChatInstance(id, 0, "")

	key := dal.NewKeyWithID(ChatInstanceCollection, id)
	tgChatInstance := TgChatInstance{
		WithID: record.NewWithID(id, key, tgChatInstanceData),
		Data:   tgChatInstanceData,
	}

	var session dal.ReadSession = tgChatInstanceDal.db
	//if tx == nil {
	//	session = tgChatInstanceDal.db
	//} else {
	//	session = tx
	//}
	if err = session.Get(c, tgChatInstance.Record); dal.IsNotFound(err) {
		return
	}
	return
}

func (tgChatInstanceDal tgChatInstanceDalgo) SaveTelegramChatInstance(c context.Context, id string, tgChatInstanceData botsfwtgmodels.TgChatInstanceData) (err error) {
	key := dal.NewKeyWithID(ChatInstanceCollection, id)
	chatInstance := record.NewWithID(id, key, tgChatInstanceData)
	err = tgChatInstanceDal.db.RunReadwriteTransaction(c, func(ctx context.Context, tx dal.ReadwriteTransaction) error {
		return tx.Set(ctx, chatInstance.Record)
	})
	return
}

func (tgChatInstanceDalgo) NewTelegramChatInstance(chatInstanceID string, chatID int64, preferredLanguage string) (tgChatInstanceData botsfwtgmodels.TgChatInstanceData) {
	tgChatInstanceData = &botsfwtgmodels.TgChatInstanceBaseData{
		TgChatID:          chatID,
		PreferredLanguage: preferredLanguage,
	}
	return tgChatInstanceData
}

const ChatInstanceCollection = "botTgChatInstance"
