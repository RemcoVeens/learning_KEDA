# learning KEDA

this is my first experience with kubernetes event-driven autoscaling. using redis in this case

`producer` creates job_queue object and waits for completion-queue object

`consumer` waits for job_queue objects, 'proccesses' it and creates completion-queue object
