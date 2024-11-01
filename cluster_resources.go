package main

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var mysqlPod = &v1.Pod{
	ObjectMeta: metav1.ObjectMeta{
		Name:   "mysql",
		Labels: map[string]string{"app": "mysql"},
	},
	Spec: v1.PodSpec{
		Containers: []v1.Container{
			{
				Name:  "mysql",
				Image: "mysql",
				Env: []v1.EnvVar{
					{
						Name:  "MYSQL_ROOT_PASSWORD",
						Value: "1234",
					},
					{
						Name:  "MYSQL_DATABASE",
						Value: "wordpress",
					},
				},
				Ports: []v1.ContainerPort{
					{
						ContainerPort: 3306,
					},
				},
			},
		},
	},
}

var mysqlService = &v1.Service{
	ObjectMeta: metav1.ObjectMeta{
		Name: "mysql",
	},
	Spec: v1.ServiceSpec{
		Selector: map[string]string{
			"app": "mysql",
		},
		Ports: []v1.ServicePort{
			{
				Protocol:   v1.ProtocolTCP,
				Port:       3306,
				TargetPort: intstr.FromInt32(3306),
			},
		},
	},
}

var wordpressPod = &v1.Pod{
	ObjectMeta: metav1.ObjectMeta{
		Name:   "wordpress",
		Labels: map[string]string{"app": "wordpress"},
	},
	Spec: v1.PodSpec{
		Containers: []v1.Container{
			{
				Name:  "wordpress",
				Image: "wordpress",
				Env: []v1.EnvVar{
					{
						Name:  "WORDPRESS_DB_HOST",
						Value: "mysql",
					},
					{
						Name:  "WORDPRESS_DB_NAME",
						Value: "wordpress",
					},
					{
						Name:  "WORDPRESS_DB_USER",
						Value: "root",
					},
					{
						Name:  "WORDPRESS_DB_PASSWORD",
						Value: "1234",
					},
				},
				Ports: []v1.ContainerPort{
					{
						ContainerPort: 80,
					},
				},
			},
		},
	},
}

var wordpressService = &v1.Service{
	ObjectMeta: metav1.ObjectMeta{
		Name: "wordpress",
	},
	Spec: v1.ServiceSpec{
		Selector: map[string]string{
			"app": "wordpress",
		},
		Ports: []v1.ServicePort{
			{
				Protocol:   v1.ProtocolTCP,
				Port:       80,
				TargetPort: intstr.FromInt32(80),
			},
		},
		Type: v1.ServiceTypeClusterIP, // Explicitly set the type
	},
}
