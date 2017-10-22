package server

import ()

var (
	TopicQueueHandler = NewTopicPool(10000)
)

func (s *Server) handleCreate(r *Request) Reply {
	var topic string

	topic = string(r.Arguments[1])

	if r.HasArgument(0) == false {
		return ErrNotEnoughArgs
	}
	TopicQueueHandler.CreateTopic(topic)

	return &StatusReply{
		code: "OK",
	}
}

func (s *Server) handleBind(r *Request) Reply {
	var topic string
	var qname string

	topic = string(r.Arguments[0])
	qname = string(r.Arguments[1])

	if r.HasArgument(0) == false {
		return ErrNotEnoughArgs
	}
	TopicQueueHandler.Bind(topic, qname)

	return &StatusReply{
		code: "OK",
	}
}

func (s *Server) handlePub(r *Request) Reply {
	var topic string
	var body string

	if r.HasArgument(0) == false {
		return ErrNotEnoughArgs
	}

	topic = string(r.Arguments[0])
	body = string(r.Arguments[1])

	TopicQueueHandler.Pub(topic, &body)

	return &StatusReply{
		code: "OK",
	}
}

func (s *Server) handleSub(r *Request) Reply {
	var topic string
	var qname string
	var body *string

	if r.HasArgument(0) == false {
		return ErrNotEnoughArgs
	}

	topic = string(r.Arguments[0])
	qname = string(r.Arguments[0])
	body = TopicQueueHandler.Sub(topic, qname)

	return &StatusReply{
		code: *body,
	}
}

func (s *Server) handleExists(r *Request) Reply {
	var id int64

	if r.HasArgument(0) == false {
		return ErrNotEnoughArgs
	}

	return &IntReply{
		number: id,
	}
}

func (s *Server) handleDel(r *Request) Reply {
	var id int64 = 0

	if r.HasArgument(0) == false {
		return ErrNotEnoughArgs
	}

	return &IntReply{
		number: id,
	}
}

func (s *Server) handleSelect(r *Request) Reply {
	if r.HasArgument(0) == false {
		return ErrNotEnoughArgs
	}

	num := string(r.Arguments[0])
	if len(num) == 0 {
		return ErrNotEnoughArgs
	}

	return &StatusReply{
		code: "OK",
	}
}
