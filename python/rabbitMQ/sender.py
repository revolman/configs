import pika


def main():
    credentials = pika.PlainCredentials('revolman', 'Shoop!r6')
    parameters = pika.ConnectionParameters(host='192.168.0.156',
                                           port=5672,
                                           virtual_host='/',
                                           credentials=credentials)
    connection = pika.BlockingConnection(parameters)
    channel = connection.channel()

    channel.queue_declare(queue="hello-python")

    while True:
        channel.basic_publish(exchange='',
                              routing_key='hello-python',
                              body=str('My mess-s-sage!'))
        print(" [x] Sent message")

    connection.close()


main()
