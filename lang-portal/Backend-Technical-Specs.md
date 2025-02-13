## Backend Server Technical Specs

## Business Goal

A language learning school wants to build a prototype of learning portal which will act as three things:
- Inventory of possible vocabulary that can be learned
- Act as a  Learning record store (LRS), providing correct and wrong score on practice vocabulary
- A unified launchpad to launch different learning apps

## Technical Requirements

- The backend will be built using Go
- The database will be SQLite3
- The API will be built using GIN
- The API will always return JSON
- There will be no authentication or authorization
- Everything will be treated as a single user

## Database Schema

Our database will be a single sqlite database called 'words.db' that will be in the root of the project folder 'lang_portal'.

We have the following tables:
- words - stored vocabulary words
    - id integer
    - spanish string 
    - english string
    - parts json
- words_groups - joint table for words and groups (many-to-many)
    - id integer
    - word_id integer
    - group_id integer
- groups - thematic groups of words
    - id integer
    - name string
    - words_count integer
- study_sessions - records of study sessions grouping word_review_items
    - id integer
    - group_id integer
    - study_activity_id
    - created_at datetime
- study_activities - a specific study activity, linking a study session to a vocabulary word
    - id integer
    - name string
    - url string
- word_review_items - a record of word practice, determining if the word was correct or not
    - id integer
    - word_id integer
    - study_session_id integer
    - correct boolean
    - created_at timestamp

## API Endpoints

### GET /api/dashboard/last_study_session
Returns information about the most recent study session.
**Response:**
```json
{
  "id": 1,
  "group_id": 1,
  "study_activity_id": 1,
  "created_at": "2023-10-01T12:00:00Z",
  "group_name": "Basic Greetings",
  "correct_count": 5,
  "wrong_count": 2
}
```

### GET /api/dashboard/study_progress
Returns study progress statistics.
**Response:**
```json
{
  "total_words_studied": 100,
  "total_available_words": 200,
  "mastery_progress": "50%"
}
```

### GET /api/dashboard/quick_stats
Returns quick statistics for the dashboard.
**Response:**
```json
{
  "success_rate": "80%",
  "total_study_sessions": 22,
  "total_active_groups": 3,
  "study_streak": "4 days"
}
```

### GET /api/study_activities/:id
Returns details of a specific study activity.
**Response:**
```json
{
  "id": 1,
  "name": "Flashcards",
  "url": "http://example.com/flashcards"
}
```

### GET /api/study_activities/:id/study_sessions
Returns a list of study sessions for a specific study activity.
**Response:**
```json
[
  {
    "id": 1,
    "group_id": 1,
    "study_activity_id": 1,
    "created_at": "2023-10-01T12:00:00Z"
  },
  {
    "id": 2,
    "group_id": 2,
    "study_activity_id": 1,
    "created_at": "2023-10-02T12:00:00Z"
  }
]
```

### POST /api/study_activities
Creates a new study activity.
**Required params:** group_id, study_activity_id
**Response:**
```json
{
  "id": 1,
  "group_id": 1,
  "study_activity_id": 1,
  "created_at": "2023-10-01T12:00:00Z"
}
```

### GET /api/words
Returns a paginated list of words.
**Response:**
```json
[
  {
    "id": 1,
    "spanish": "hola",
    "english": "hello",
    "parts": ["greeting"]
  },
  {
    "id": 2,
    "spanish": "adios",
    "english": "goodbye",
    "parts": ["farewell"]
  }
]
```

### GET /api/words/:id
Returns details of a specific word.
**Response:**
```json
{
  "id": 1,
  "spanish": "hola",
  "english": "hello",
  "parts": ["greeting"],
  "correct_count": 10,
  "wrong_count": 2
}
```

### GET /api/groups
Returns a list of groups.
**Response:**
```json
[
  {
    "id": 1,
    "name": "Greetings",
    "words_count": 10
  },
  {
    "id": 2,
    "name": "Farewells",
    "words_count": 5
  }
]
```

### GET /api/groups/:id
Returns details of a specific group.
**Response:**
```json
{
  "id": 1,
  "name": "Greetings",
  "stats": {
    "words_count": 10
  }
}
```

### GET /api/groups/:id/words
Returns a list of words in a specific group.
**Response:**
```json
[
  {
    "id": 1,
    "spanish": "hola",
    "english": "hello",
    "correct_count": 5,
    "wrong_count": 2
  },
  {
    "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 20,
    "items_per_page": 100
    }
    }
]
```

### GET /api/groups/:id/study_sessions
Returns a list of study sessions for a specific group.
**Response:**
```json
[
  {
    "id": 1,
    "group_id": 1,
    "study_activity_id": 1,
    "created_at": "2023-10-01T12:00:00Z"
  },
  {
    "id": 2,
    "group_id": 1,
    "study_activity_id": 2,
    "created_at": "2023-10-02T12:00:00Z"
  }
]
```

### GET /api/study_sessions/:id
Returns details of a specific study session.
**Response:**
```json
{
  "id": 1,
  "group_id": 1,
  "study_activity_id": 1,
  "created_at": "2023-10-01T12:00:00Z",
  "activity_name": "Flashcards",
  "group_name": "Greetings",
  "start_time": "2023-10-01T12:00:00Z",
  "end_time": "2023-10-01T12:30:00Z",
  "number_of_review_items": 10
}
```

### GET /api/study_sessions/:id/words
Returns a list of words reviewed in a specific study session.
**Response:**
```json
[
  {
    "id": 1,
    "word_id": 1,
    "study_session_id": 1,
    "correct": true,
    "created_at": "2023-10-01T12:00:00Z"
  },
  {
    "id": 2,
    "word_id": 2,
    "study_session_id": 1,
    "correct": false,
    "created_at": "2023-10-01T12:05:00Z"
  }
]
```

### POST /api/reset_history
Resets the study history.
**Response:**
```json
{
  "status": "success"
}
```

### POST /api/full_reset
Performs a full reset of the database.
**Response:**
```json
{
  "status": "success"
}
```

### POST /api/study_sessions/:id/words/:word_id/review
Records a review of a word in a study session.
**Required params:** correct
**Response:**
```json
{
  "id": 1,
  "word_id": 1,
  "study_session_id": 1,
  "correct": true,
  "created_at": "2023-10-01T12:00:00Z"
}
```
## Mage tasks

Mage is a task runner for go. Let's list out possible tasks we need for our lang portal.

### Initialize Database
This task will initialize the SQLite called 'words.db'.

### Migrate database
This task will run a series of migraitons sql files on the database.

Migrations live in the `migrations` folder.
The migration files will be run in order of their file name.
The file names should look like this:

````sql
0001_init.sql
0002_create_words_table.sql
````

### Seed Data
This task will import json files and transform them into target data for our database.

All seed files live in the 'seeds' folder.
All seed files should be loaded.
In our task we should have a DSL to specify each seed file and its expected group word name.

````json
[
    {
        "spanish": "hola",
        "english": "hello"
    },
]
````