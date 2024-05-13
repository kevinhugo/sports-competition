# sports-competition

Sport Competition project for VISKA Technical Assessment

## Basic Installation

1. Copy the .env.example and rename it to .env and adjust the values inside .env
2. run docker-compose up --build -d
3. to use API via Swagger, please go to {{base_url}}/sports-competition/swagger/index.html ( example : http://127.0.0.1:9000/sports-competition/swagger/index.html# )

## Database

Database migration will be handled automatically by Gorm during service start, but if eagerly wants to do it manually, database dumpfile schema is provided at app/database/sports_competition_dumpfile.sql

## Unit Tests

Unit test will automatically run while building Docker at layer 10, if unit test fails, it will abort the building process of the Docker, stopping the docker-compose process. ( unit_test.go files were meant to be separated for modular purpose but current technical difficulties made it unavailable for the moment )

## Step-by-step API Usage

### First Step
Firstly, access the Login API to get Authorization Token [URL : {{base_url}}/sports-competition/v1/user/login]

API accepts username and password, if username is unique, the provided data will be considered as registered data to be registered, if username already exists, valid password must be given or else it will return invalid password.

### Second Step
Secondly, after you've received the access token you can start using the Begin Competition API [URL : {{base_url}}/sports-competition/v1/sport/begin-competition]

Must be remembered that you need to put the token to the Authorization header or else it'll return 401.

Example Data :

{

    "base_skill" : [5,9],

    "opponent_proficiencies" : [2,3,6,7,8],

    "opponent_exps" : [3,4,2,2,3]

}

Validations ( if validity check is failed, process will abort ):
* base_skill : index 0 : length of opponent_proficiencies and opponent_exps, index 1 : base proficiency value
* opponent_proficiencies : list of opponent proficiencies ( note that the length must match the index 0 of base_skill )
* opponent_exps : list of opponent gained proficiency points if defeated ( note that the length must match the index 0 of base_skill )


### Third Step
Thirdly, access the Get Identity API [URL : {{base_url}}/sports-competition/v1/user/identity/{{first_name}}]

first_name servers as param that will be shown on response 