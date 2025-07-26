package di

import (
	"post-tech-challenge-10soat/internal/controllers"
	"post-tech-challenge-10soat/internal/delivery/http/handler"
	"post-tech-challenge-10soat/internal/external/mongo"
	repositorymongo "post-tech-challenge-10soat/internal/external/mongo/repositorymongo"
	"post-tech-challenge-10soat/internal/external/postgres"
	repository "post-tech-challenge-10soat/internal/external/postgres/repositories"
	"post-tech-challenge-10soat/internal/gateways"
	"post-tech-challenge-10soat/internal/infrastructure/config"
	"post-tech-challenge-10soat/internal/infrastructure/logger"
	"post-tech-challenge-10soat/internal/usecases/client"
	"post-tech-challenge-10soat/internal/usecases/order"
	"post-tech-challenge-10soat/internal/usecases/product"
)

func Setup(config *config.App, db *postgres.DB, mongo *mongo.MONGO) (
	handler.HealthHandler,
	handler.ClientHandler,
	handler.ProductHandler,
	handler.OrderHandler) {
	logger.Set(config)

	// Repositories
	clientRepo := repositorymongo.NewClientMongoRepositoryImpl(mongo.Database)
	productRepo := repository.NewProductRepositoryImpl(db)
	categoryRepo := repository.NewCategoryRepositoryImpl(db)
	orderRepo := repository.NewOrderRepositoryImpl(db)
	orderProductRepo := repository.NewOrderProductRepositoryImpl(db)
	// paymentRepo := repository.NewPaymentRepositoryImpl(db)

	// Gateways
	clientGateway := gateways.NewClientGatewayImpl(
		clientRepo,
	)
	productGateway := gateways.NewProductGatewayImpl(
		productRepo,
	)
	categoryGateway := gateways.NewCategoryGatewayImpl(
		categoryRepo,
	)
	orderGateway := gateways.NewOrderGatewayImpl(
		orderRepo,
	)
	orderProductGateway := gateways.NewOrderProductGatewayImpl(
		orderProductRepo,
	)
	// paymentGateway := gateways.NewPaymentGatewayImpl(
	// 	paymentRepo,
	// )

	// Usecases
	getClientByCpf := client.NewGetClientByCpfUseCaseImpl(
		clientGateway,
	)
	getClientById := client.NewGetClientByCpfUseCaseImpl(
		clientGateway,
	)
	createClient := client.NewCreateClientUsecaseImpl(
		clientGateway,
	)
	createProduct := product.NewCreateProductUsecaseImpl(
		productGateway,
		categoryGateway,
	)
	updateProduct := product.NewUpdateProductUsecaseImpl(
		productGateway,
		categoryGateway,
	)
	deleteProduct := product.NewDeleteProductUsecaseImpl(
		productGateway,
	)
	listProducts := product.NewListProductsUsecaseImpl(
		productGateway,
		categoryGateway,
	)
	// paymentUseCase := payment.NewPaymentCheckoutUsecaseImpl(
	// 	paymentGateway,
	// )
	createOrder := order.NewCreateOrderUsecaseImpl(
		productGateway,
		clientGateway,
		orderGateway,
		orderProductGateway,
	)
	listOrders := order.NewListOrdersUseCaseImpl(
		orderGateway,
	)
	getOrderPaymentStatus := order.NewGetOrderPaymentStatusUseCaseImpl(
		orderGateway,
	)
	updateOrderStatus := order.NewUpdateOrderStatusUseCaseImpl(
		orderGateway,
	)

	// Controllers
	clientController := controllers.NewClientController(
		getClientByCpf,
		getClientById,
		createClient,
	)
	productController := controllers.NewProductController(
		createProduct,
		deleteProduct,
		updateProduct,
		listProducts,
	)
	orderController := controllers.NewOrderController(
		createOrder,
		listOrders,
		getOrderPaymentStatus,
		updateOrderStatus,
	)

	// Handlers
	healthHandler := handler.NewHealthHandler()
	clientHandler := handler.NewClientHandler(clientController)
	productHandler := handler.NewProductHandler(*productController)
	orderHandler := handler.NewOrderHandler(*orderController)

	return healthHandler, clientHandler, productHandler, orderHandler
}
