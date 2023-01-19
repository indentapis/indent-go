/**

Auth is currently handled with the `indent` cli.
$ indent auth login --space my-space-name

To verify you have authenticated successfully, you can check `auth view`:
$ indent auth view --space my-space-name 2>&1 | jq
{
  ...
  "msg": "Current User",
  "user": {
    "Id": "my-email@example.com",
    ...
    "Labels": [
        "indent.com/app/config/id": "space:my-space-name"
    ]
    ...
  }
}
**/

provider "indent" {
  space = "my-space-name"
}
