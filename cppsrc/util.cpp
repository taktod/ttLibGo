//#include "ttLibC/ttLibC/util/stlMapUtil.h"
#include <string>
#include <map>
#include "util.hpp"

using namespace std;

string maps::getString(string key) {
  auto iter = strmap.find(key);
  if(iter != strmap.end()) {
    return iter->second;
  }
  return "";
}
uint32_t maps::getUint32(string key) {
  auto iter = uint32map.find(key);
  if(iter != uint32map.end()) {
    return iter->second;
  }
  return 0;
}
uint64_t maps::getUint64(string key) {
  auto iter = uint64map.find(key);
  if(iter != uint64map.end()) {
    return iter->second;
  }
  return 0;
}
list<string> maps::getStringList(string key) {
  auto iter = strlistmap.find(key);
  if(iter != strlistmap.end()) {
    return iter->second;
  }
  list<string> empty;
  return empty;
}

extern "C" {

// string -> stringで持っておいて、数値だったら、castするか・・・
void *StdMap_make() {
  maps *m = new maps();
  return (void *)m;
}
void StdMap_putString(void *ptr, const char *key, const char *value) {
  maps *m = reinterpret_cast<maps *>(ptr); 
  m->strmap.insert(pair<string, string>(key, value));
}
void StdMap_putUint32(void *ptr, const char *key, uint32_t value) {
  maps *m = reinterpret_cast<maps *>(ptr); 
  m->uint32map.insert(pair<string, uint32_t>(key, value));
}
void StdMap_putUint64(void *ptr, const char *key, uint64_t value) {
  maps *m = reinterpret_cast<maps *>(ptr); 
  m->uint64map.insert(pair<string, uint32_t>(key, value));
}

void StdMap_putStringList(void *ptr, const char *key, const char *value) {
  maps *m = reinterpret_cast<maps *>(ptr); 
  auto iter = m->strlistmap.find(key);
  if(iter == m->strlistmap.end()) {
    // データがないので、作らなければならない。
    list<string> newList;
    newList.push_back(value);
    m->strlistmap.insert(pair<string, list<string>>(key, newList));
  }
  else {
    list<string> l = iter->second;
    l.push_back(value);
    m->strlistmap.erase(key);
    m->strlistmap.insert(pair<string, list<string>>(key, l));
  }
}

void StdMap_close(void *ptr) {
  maps *m = reinterpret_cast<maps *>(ptr); 
  delete m;
}

}
