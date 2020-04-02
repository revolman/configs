import pika
import json
import time


def callback(ch, method, properties, body):
    print(" [x] Получено сообщение: %r" % body)

    message = json.loads(body)

    result = message["first_arg"] + message["second_arg"]

    print(" [x] Результат: %s + %s = %s" % (message["first_arg"], message["second_arg"], result))

    time.sleep(1)

    print(" [x] Готово")
    ch.basic_ack(delivery_tag=method.delivery_tag)


if __name__ == '__main__':
    credentials = pika.PlainCredentials('USERNAME', 'PASSWD')
    parameters = pika.ConnectionParameters(host='192.168.0.156',
                                           port=5672,
                                           virtual_host='/',
                                           credentials=credentials)
    connection = pika.BlockingConnection(parameters)

    channel = connection.channel()

    channel.queue_declare(queue="worker-q1-python", durable=True)
    print(' [*] Ожидаю сообщения. Для выхода нажмите CTRL+C.')

    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(queue='worker-q1-python', on_message_callback=callback)

    channel.start_consuming()
