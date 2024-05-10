## Get Ready to Code!

We're excited to see your programming skills in action! This is a relaxed environment where you can showcase your abilities and build a fantastic TODO application using your preferred technology stack.

**Objectives:**

-   Develop a TODO application using your familiar tools.

**Data Requirements:**

The application will manage tasks with the following attributes:

| Field  | Data Type | Notes |
| ------ | --------- | ----- |
|ID|Any (Unique)|Must be a unique identifier for each task.|
|Title |String (100 char) |The title of the task, limited to 100 characters.|
|Description|String|A detailed description of the task.|
|CreatedAtDate/Time (UTC)|The date and time the task was created, stored in UTC format.|
|UpdatedAt|Date/Time (UTC)|The date and time the task was last updated, stored in UTC format.|
|Image|String (Base64)|An optional image encoded in Base64 format.|
|Status|String (Enum)|The current status of the task, must be either "IN_PROGRESS" or "COMPLETED".|

**Functionality:**

-   **Create Tasks:**
    
    -   Users can create new tasks, specifying the `Title`, `Date` (RFC3399 with Timezone format), and `Status` (`IN_PROGRESS` or` COMPLETED`).
    -   The `ID` will be automatically generated and must be unique.
    -   The `Title` cannot exceed 100 characters.
    -   The `Image` field can optionally store a Base64 encoded image.
    
-   **List Tasks:**
    
    -   Retrieve a list of all tasks.
    -   Sort tasks by` Title`, `Date`, or `Status`.
    -   Search tasks by `Title` or `Description`.
    
-   **Update Tasks:**
    
    -   Update existing tasks by their `ID`.
    -   Update any of the `Title`, `Description`,` Date`, `Image`, or `Status` fields with the same requirements as when creating a task.
    

**Bonus Points:**

Feel free to add any creative features that showcase your programming skills and impress us!

**Application Specs:**

-   **API Framework:**
    
    -   Use any API framework you're comfortable with.
    -   Bonus points for using `Golang` with `gRPC.`
    
-   **Data Format:**
    
    -   Read and write data in `JSON` format.
    
-   **API Design:**
    
    -   Implement a `RESTful API` design.
    
-   **Testing:**
    
    -   Write unit tests for your code.
    
-   **Bonus Features:**
    -   Develop the application using `Hexagonal Architecture` principles.
    -   Create an `API Explorer` using `Swagger`.
    -   Package your application as a `Docker image`.