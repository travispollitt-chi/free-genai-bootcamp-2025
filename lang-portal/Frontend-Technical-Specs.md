## Frontend Technical Spec

## Pages

## Dashboard `/dashboard`

## Purpose
The purpose of this page is to provide a summary of learning and act as the default page when a user is visiting the web app.

## Components
This page contains the following components
- Last Study Session
    - shows last activity used
    - shows when last activity used
    - summarizes wrong vs correct from last activity
    - has a link to the group

- Study Progress
    - total words studied
        - across all study sessions show the total words studied - out of all possible words in our database
    - display a mastery of progress e.g. 0%

- Quick Stats
    - sucess rate eg. 80%
    - total study sessions eg. 22
    - total active groups eg. 3
    - study streak eg. 4 days

- Start Studying Button
    - goes to study activities 

We'll need following API endpoints to power this page
- GET /dashboard/last_study_session
- GET /dashboard/study_progress
- GET /dashboard/quick stats

### Study Activities `/study-activities`

## Purpose
The purpose of this page is to show a collection of study activities with a thumbnail and it's name, to either launch or view the study activity.

## Components

- Study activity card
    - Show a thumbnail of the study activity
    - the name of the study activity
    - a launch button to take us to the launch page
    - the view page to view more information about past study sessions for this study activity

## Needed API Endpoints

- GET /study_activities/:id
- GET /study_activities/:id/study_sessions

### Study Activity Show `/study_activities/:id`

## Purpose
The purpose of this page is to show the details of a study activity and its past study sessions.

## Components
- Name of study activity
- Thumbnail of study activity
- Description of study activity
- Launch button
- Study Activities Paginated List
    - id
    - activity name
    - group name
    - start time
    - end time (inferred by the last word_review_item submitted)
    - number of review items

## Needed API Endpoints
- GET /api/study_activities/:id
- GET /api/study_activities/:id/study_sessions

### Study Activities Launch `/study_activities/:id/launch`

## Purpose
The purpose of this page is to launch a study activity.

## Components
- Name of study activity
- Launch form
    - select field for group
    - launch now button

## Behavior
After the form is submitted a new tab opens with the study activity based on its URL provided in the database.

Also after the form is submmitted the page will redirect to the study session showpage.

## Needed API Endpoints
- POST /api/study_activities
    - Required 

## Words `/words`

## Purpose
The purpose of this page is to show all words in our database.

## Components
- Paginated Word List
    - Columns
        - Spanish
        - Englsih
        - Correct Count
        - Wrong Count
    - Pagination with 100 items per page
    - Clicking the Spanish word will take us to the word show page 

## Needed API Endpoints
- GET /api/words

### Word Show `/words/:id`

## Purpose
The purpose of this page is to show information about a specific word.

## Components
- Spanish
- English
- Study Statistics
    - Correct Count
    - Wrong Count
- Word Groups
    - Show on a series of pills eg. tags
    - When group name is clicked it will take us to the group show page

## Needed API Endpoints
- GET /api/words/:id

### Word Groups `/groups`

## Purpose
The purpose of this page is to show a list of groups in our database.

## Components
- Paginated Group List
    - Columns
        - Group Name
        - Word Count
    - Clicking the group name will take us to the group show page

## Needed API Endpoints
- GET /api/groups

### Group Show `/groups/:id`

## Purpose
The purpose of this page is to show information about a specific group.

## Components
- Group Name
- Group Statistics
    - Total Word Count
- Words in Group (Paginated List of Words)
    - Should use the same component
- Study Sessions (Paginated List of Study Sessions)
    - Should use the same component as the study sessions index page

## Needed API Endpoints
- GET /api/groups/:id (the name and groups stats)
- GET /api/groups/:id/words
- GET /api/groups/:id/study_sessions


----
## Purpose
## Components
## Needed API Endpoints