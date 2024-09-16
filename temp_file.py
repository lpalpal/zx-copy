def greet(name):
  """
  This function greets the person passed in as a parameter.

  Args:
    name: The name of the person to greet.

  Returns:
    A greeting string.
  """
  return f"Hello, {name}!"

if __name__ == "__main__":
  """
  This code will only run when the script is executed directly.
  """
  user_name = input("Enter your name: ")
  print(greet(user_name))
