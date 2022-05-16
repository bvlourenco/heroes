class ServiceException {
  final String message;

  ServiceException(this.message);

  @override
  String toString() => message;
}
