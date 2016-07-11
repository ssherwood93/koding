amqp     = require 'amqp'
KONFIG   = require 'koding-config-manager'
Tracker  = require '../../../social/lib/social/models/tracker.coffee'

sendmail = (require 'sendmail') {
  logger   : {  # disable verbosed loggers
    debug  : ->
    info   : ->
    warn   : ->
    error  : ->
  }
}

module.exports = class EmailerWorker

  TEMPLATES  = require './templates'
  MAILEVENTS = [
    'REQUEST_NEW_PASSWORD'
    'CHANGED_PASSWORD'
    'REQUEST_EMAIL_CHANGE'
    'CHANGED_EMAIL'
    'INVITED_TEAM'
    'REQUESTED_TEAM_LIST'
  ]

  getEmailType = (subject) ->
    return type  for type, val of Tracker.types when val is subject

  isEmailEvent = (subject) ->
    return (getEmailType subject) in MAILEVENTS

  ['log', 'error'].forEach (logger) ->
    EmailerWorker::[logger] = (message...) ->
      console[logger] '[EmailerWorker]', message...

  QUEUENAME    = 'NodeMailSender:0:WorkerQueue'
  QUEUEOPTIONS = { durable: yes, autoDelete: no }

  EXCHANGENAME = 'BrokerMessageBus:0'
  EMAILEVENT   = 'api.mail_send'


  start: ->

    @log 'starting...'

    @connection = amqp.createConnection KONFIG.mq, { reconnect: yes }

    @connection.on 'error', (err) =>
      @error 'Error: connecting to RabbitMQ', err

    @connection.on 'ready', =>
      @log 'started successfully'
      @createQueue()


  createQueue: ->

    @connection.queue QUEUENAME, QUEUEOPTIONS, (queue) =>

      @log 'queue created'

      queue.bind EXCHANGENAME, '#', =>

        queue.subscribe (message, header, property) =>

          return  unless property.type is EMAILEVENT

          message = JSON.parse message.data.toString()

          return  unless isEmailEvent message.Subject

          @sendMail message


  sendMail: (message) ->

    type = getEmailType message.Subject

    message.Properties.Options.username = message.Properties.Username
    mail = (TEMPLATES message.Properties.Options)[type]

    @log "sending #{type} mail to #{message.To}..."

    sendmail

      from    : 'hello@koding.com'
      to      : message.To
      subject : mail.subject
      content : mail.content

    , (err) =>

      if err
      then @error "failed to send #{type} mail to #{message.To}", err
      else @log "successfully sent #{type} mail to #{message.To}"
