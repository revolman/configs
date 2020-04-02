import pika
import random
import json


if __name__ == '__main__':
    credentials = pika.PlainCredentials('USERNAME', 'PASSWD')
    parameters = pika.ConnectionParameters(host='192.168.0.156',
                                           port=5672,
                                           virtual_host='/',
                                           credentials=credentials)
    connection = pika.BlockingConnection(parameters)

    channel = connection.channel()

    channel.exchange_declare(exchange='workers', exchange_type='direct')
    channel.queue_declare(queue='worker-q1-python', durable=True)
    channel.queue_bind(exchange='workers', queue='worker-q1-python')

    index = 0
    while True:
        index += 1
        if index > 1000:
            break

        message = {
            "id": index,
            "first_arg": random.randint(1, 100),
            "second_arg": random.randint(1, 100),
        }

        channel.basic_publish(
            exchange='workers',
            routing_key='worker-q1-python',
            body=json.dumps(message),
            properties=pika.BasicProperties(
                delivery_mode=2
            ))
        print(" [x] Sent %r" % message)

    connection.close()
