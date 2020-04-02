import pika


def callback(ch, method, properties, body):
    print(" [x] Получено сообщение: %r" % body)


if __name__ == '__main__':
    credentials = pika.PlainCredentials('revolman', 'Shoop!r6')
    parameters = pika.ConnectionParameters(host='192.168.0.156',
                                           port=5672,
                                           virtual_host='/',
                                           credentials=credentials)
    connection = pika.BlockingConnection(parameters)
    channel = connection.channel()

    channel.queue_declare(queue="hello-python")

    channel.basic_consume(queue='hello-python',
                          auto_ack=True,
                          on_message_callback=callback)
    print(' [*] Ожидаю сообщения. Для выхода нажмите CTRL+C.')
    channel.start_consuming()
