# 虚拟机/物理机占用CPU、Mem小工具

## 功能
- 占用cpu利用率40%以上
- 占用mem百分之60%以上，过高后会自动GC，释放内存

## crontab

建议使用crontab定时任务，避免重启或者其它误杀情况

```shell
crontab -e
*/30 * * * * /root/sttool.sh
```
### 欢迎提PR，共同完善