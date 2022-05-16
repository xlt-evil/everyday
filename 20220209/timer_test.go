package _0220209

/**
 * @Author: hxy
 * @Description:
 * @File:  timer_test
 * @Date: 2022/02/09 11:34
 */

// timer 使用的坑
// 错误的创建多个timer ，不能再循环里面调用timer ,因为timer 是会被提升到全局去的，通过GMP调用，应该使用reset 重新设置 ，使用STOP 时需要手动释放channel