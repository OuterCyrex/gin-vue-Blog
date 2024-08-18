# gin-vue-blog

#### Description

I am pleased to announce the official release of version 1.0 on August 18, 2024
This project was initiated on August 6, 2024 and took a total of 13 days to develop During this period, I encountered many difficulties, such as the outdated JWT library, and at the beginning When writing the frontend, the UI libraries are not compatible with Vue2 or something like that.
But in the end, they were all successfully overcome, and the result was good!

## introduce

For an attempt to separate front-end and back-end projects, the front-end was built using the `Vue2` framework, the UI library used `ant design vue` and `Vuetify`, and the back-end used frameworks such as `gin` and `gorm` in the Golang
Overall, the effect of logging in, commenting, and viewing articles has been achieved.
Due to the testing being conducted on the computer, the bugs that exist when opening the website on the mobile end are still unknown and need to be tested after going online.

## Q＆A
#### How to achieve login?
There is a gray avatar box in the top left corner of the front-end display page, click to enter the login interface.
The 1.0 version currently does not support custom avatars because the online space of the image bed is limited (although considering that there won't be many users, it doesn't really matter).
And in order to prevent malicious registration and other network attacks, version 1.0 temporarily uses an invitation system to register new accounts, which means that as long as you tell the webmaster, they will be assigned different accounts.
![](http://shwfbbqxt.hd-bkt.clouddn.com/6H3VF3%5D5Y%5DXX%25~24N%60P%7B%29VN.png)

#### What is the purpose of logging in?
Actually, it's not very useful, mainly for making comments. In order to improve the quality of comments and limit malicious comments, only logged in users can make comments.
However, after all, a personal blog is a personalized website, so its main function is still to display articles. Therefore, whether logged in or not, you can access the content of articles.
![](http://shwfbbqxt.hd-bkt.clouddn.com/%28%24WDGZ%28B%7B8M1QP%28Q%5D_U6QS2.png)

#### What happens if you comment on something bad?
It's not really that bad, it's just that if I see inappropriate comments online, I will naturally delete them in the background.
Being crazy or anything is allowed, but things that are too abnormal are not allowed!

------

### Project reference materials:
Wenjue Chen's blog project: https://github.com/wejectchen/Ginblog
