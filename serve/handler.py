def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """

    bytes = None
    with open("./function/index.html") as f:
        bytes = f.read()

    return str(bytes)
