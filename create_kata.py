main_file_template = """package main

func FUNCTION_NAME(ARGUMENTS) RETURN_TYPE {
	var response RETURN_TYPE

	return response
}
"""

test_file_template = """package main

import "testing"

func Test_FUNCTION_NAME(t *testing.T) {
	type args struct {
		ARGUMENTS
	}

	type test struct {
		name string
		args args
		want RETURN_TYPE
	}

	tests := []test{
		{
			name: "#1",
			args: args{
				ARGUMENTS
			},
			want: RETURN_TYPE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FUNCTION_NAME(tt.args.input); got != tt.want {
				t.Errorf("FUNCTION_NAME() %s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
"""

# ask user for FUNCTION_NAME, ARGUMENTS, RETURN_TYPE
function_name = input("Function name: ")
arguments = input("Arguments: ")
return_type = input("Return type: ")

def pascal_case_to_snake_case(functionName):
    snake_case = ""
    for i in range(len(functionName)):
        if functionName[i].isupper():
            snake_case += "_"
        snake_case += functionName[i].lower()
    return snake_case[1:]

file_name = pascal_case_to_snake_case(function_name)

print("Creating file: ", file_name + ".go")

main_file = open("katas/" + file_name + ".go", "w")
main_file.write(main_file_template.replace("FUNCTION_NAME", function_name).replace("ARGUMENTS", arguments).replace("RETURN_TYPE", return_type))
main_file.close()

print("Creating file: ", file_name + "_test" + ".go")

test_file = open("katas/" + file_name + "_test" + ".go", "w")
test_file.write(test_file_template.replace("FUNCTION_NAME", function_name).replace("ARGUMENTS", arguments).replace("RETURN_TYPE", return_type))
test_file.close()

print("Done!")