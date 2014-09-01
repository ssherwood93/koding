package models

import (
	"socialapi/workers/common/runner"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChannelMessageListFetchMessageChannels(t *testing.T) {
	r := runner.New("test")
	if err := r.Init(); err != nil {
		t.Fatalf("couldnt start bongo %s", err.Error())
	}
	defer r.Close()

	Convey("while fethcing channel message of a message", t, func() {
		Convey("channels should be valid", func() {
			c1 := createNewChannelWithTest()
			So(c1.Create(), ShouldBeNil)

			c2 := createNewChannelWithTest()
			So(c2.Create(), ShouldBeNil)

			c3 := createNewChannelWithTest()
			So(c3.Create(), ShouldBeNil)

			cm := NewChannelMessage()
			cm.Body = "gel beri abi"
			cm.AccountId = c1.CreatorId
			cm.InitialChannelId = c1.Id
			So(cm.Create(), ShouldBeNil)

			// add to first channel
			_, err := c1.AddMessage(cm.Id)
			So(err, ShouldBeNil)

			// add to second channel
			_, err = c2.AddMessage(cm.Id)
			So(err, ShouldBeNil)

			// add to 3rd channel
			_, err = c3.AddMessage(cm.Id)
			So(err, ShouldBeNil)

			channels, err := NewChannelMessageList().FetchMessageChannels(cm.Id)
			So(err, ShouldBeNil)
			So(len(channels), ShouldEqual, 3)

			So(c1.Name, ShouldEqual, channels[0].Name)
			So(c2.Name, ShouldEqual, channels[1].Name)
			So(c3.Name, ShouldEqual, channels[2].Name)

			So(c1.GroupName, ShouldEqual, channels[0].GroupName)
			So(c2.GroupName, ShouldEqual, channels[1].GroupName)
			So(c3.GroupName, ShouldEqual, channels[2].GroupName)
		})
	})
}

func TestChannelMessageListCount(t *testing.T) {
	r := runner.New("test")
	if err := r.Init(); err != nil {
		t.Fatalf("couldnt start bongo %s", err.Error())
	}
	defer r.Close()

	Convey("while counting messages", t, func() {
		Convey("it should error if channel id is not set", func() {
			cml := NewChannelMessageList()

			c, err := cml.Count(0)
			So(err, ShouldNotBeNil)
			So(err, ShouldEqual, ErrChannelIdIsNotSet)
			So(c, ShouldEqual, 0)
		})

		Convey("it should not error if message in the channel", func() {
			// create message
			cm := createMessageWithTest()
			So(cm.Create(), ShouldBeNil)

			c := createNewChannelWithTest()
			So(c.Create(), ShouldBeNil)

			_, err := c.AddMessage(cm.Id)
			So(err, ShouldBeNil)

			cml := NewChannelMessageList()
			cml.ChannelId = c.Id

			cnt, err := cml.Count(cml.ChannelId)
			So(err, ShouldBeNil)
			So(cnt, ShouldEqual, 1)
		})

		Convey("it should not count message if account of message is troll ", func() {
			// create channel
			c := createNewChannelWithTest()
			So(c.Create(), ShouldBeNil)

			// create account as troll
			acc1 := createAccountWithTest()
			acc1.IsTroll = true
			So(acc1.Create(), ShouldBeNil)

			acc2 := createAccountWithTest()
			So(acc2.Create(), ShouldBeNil)

			// create message that creator is troll
			msg := createMessageWithTest()
			msg.AccountId = acc1.Id
			So(msg.Create(), ShouldBeNil)

			msg2 := createMessageWithTest()
			msg2.AccountId = acc2.Id
			So(msg2.Create(), ShouldBeNil)

			// add message to the channel
			// account is troll !!
			_, err := c.AddMessage(msg.Id)
			So(err, ShouldBeNil)

			_, erre := c.AddMessage(msg2.Id)
			So(erre, ShouldBeNil)

			cml := NewChannelMessageList()
			cml.ChannelId = c.Id

			// there is 2 message in the channel
			// but one of the account of message is troll so;
			// Count will not count this message
			// messages of troll is not valid to count
			cnt, err := cml.Count(cml.ChannelId)
			So(err, ShouldBeNil)
			So(cnt, ShouldEqual, 1)
		})
	})
}

func TestChannelMessageListisExempt(t *testing.T) {
	r := runner.New("test")
	if err := r.Init(); err != nil {
		t.Fatalf("couldnt start bongo %s", err.Error())
	}
	defer r.Close()

	Convey("while testing is exempt", t, func() {
		Convey("it should have error if message id is not set ", func() {
			cml := NewChannelMessageList()

			is, err := cml.isExempt()
			So(err, ShouldNotBeNil)
			So(err, ShouldEqual, ErrMessageIdIsNotSet)
			So(is, ShouldEqual, false)
		})

		Convey("it should return true is channel is exempt", func() {
			// create account as troll
			acc := createAccountWithTest()
			acc.IsTroll = true
			So(acc.Create(), ShouldBeNil)
			// create channel
			c := createNewChannelWithTest()
			c.CreatorId = acc.Id

			// create message
			msg := createMessageWithTest()
			msg.AccountId = acc.Id
			So(msg.Create(), ShouldBeNil)

			cml := NewChannelMessageList()
			cml.ChannelId = c.Id
			cml.MessageId = msg.Id

			is, err := cml.isExempt()
			So(err, ShouldBeNil)
			So(is, ShouldEqual, true)
		})

		Convey("it should return false is channel is not exempt", func() {
			// create account as not troll
			acc := createAccountWithTest()
			acc.IsTroll = false
			So(acc.Create(), ShouldBeNil)
			// create channel
			c := createNewChannelWithTest()
			c.CreatorId = acc.Id

			// create message
			msg := createMessageWithTest()
			msg.AccountId = acc.Id
			So(msg.Create(), ShouldBeNil)

			cml := NewChannelMessageList()
			cml.ChannelId = c.Id
			cml.MessageId = msg.Id

			is, err := cml.isExempt()
			So(err, ShouldBeNil)
			So(is, ShouldEqual, false)
		})
	})
}

func TestChannelMessageListMarkIfExempt(t *testing.T) {
	r := runner.New("test")
	if err := r.Init(); err != nil {
		t.Fatalf("couldnt start bongo %s", err.Error())
	}
	defer r.Close()

	Convey("while marking if channel is exempt", t, func() {
		Convey("it should have error if message id is not set", func() {
			cml := NewChannelMessageList()

			err := cml.MarkIfExempt()
			So(err, ShouldNotBeNil)
			So(err, ShouldEqual, ErrMessageIdIsNotSet)
		})

	})
}
